//go:build !frontend
// +build !frontend

package emailprotocol

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-message/mail"
	"github.com/robin-samuel/mailstream"
	"github.com/teivah/broadcast"
)

type MailClient struct {
	*broadcast.Relay[*mailstream.Mail]

	ctx context.Context

	client   *imapclient.Client
	email    string
	password string
	lastUID  imap.UID
	mutex    sync.Mutex
}

func getLatestUid(c *imapclient.Client) (imap.UID, error) {
	// Get mailbox status with UIDNEXT to find the highest UID
	status, err := c.Status("INBOX", &imap.StatusOptions{
		UIDNext: true,
	}).Wait()
	if err != nil {
		return 0, err
	}

	if status.UIDNext == 1 {
		// No messages in the mailbox
		return 0, nil
	}

	// The latest UID will be UIDNEXT - 1
	return status.UIDNext - 1, nil
}

func newMailClient(ctx context.Context, host string, port int, email, password string) (*MailClient, error) {
	// Extract the domain from the email
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid email format")
	}

	// Connect to the IMAP server
	c, err := imapclient.DialTLS(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		return nil, err
	}

	// Login to the IMAP server
	if err := c.Login(email, password).Wait(); err != nil {
		return nil, err
	}

	// Select the INBOX
	if _, err := c.Select("INBOX", nil).Wait(); err != nil {
		return nil, err
	}

	// Get the latest uid
	latestUID, err := getLatestUid(c)
	if err != nil {
		return nil, fmt.Errorf("error getting latest UID: %w", err)
	}

	// Initialize the broadcaster
	fmt.Println("Initializing broadcaster...")
	broadcaster := broadcast.NewRelay[*mailstream.Mail]()

	mc := &MailClient{
		ctx:   ctx,
		Relay: broadcaster,

		client:   c,
		email:    email,
		password: password,
		lastUID:  latestUID,
	}

	// Start polling for new messages
	go mc.poll()

	return mc, nil
}

func (mc *MailClient) poll() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Checking for new messages...")
			mc.checkForNewMessages()
			fmt.Println("Done.")
		case <-mc.ctx.Done():
			fmt.Println("Context done. Closing imapclient...")
			mc.client.Logout()
			return
		}
	}
}

func (mc *MailClient) checkForNewMessages() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	newLatest, err := getLatestUid(mc.client)
	if err != nil {
		log.Println("Error getting latest uid:", err)
		return
	} else if newLatest == mc.lastUID {
		// No new messages
		return
	}

	// Search for messages with UID greater than lastUID
	var uidSet imap.UIDSet
	uidSet.AddRange(mc.lastUID+1, newLatest)

	fmt.Printf("Searching for %d new messages...\n", newLatest-mc.lastUID)
	s, err := mc.client.UIDSearch(&imap.SearchCriteria{
		UID: []imap.UIDSet{uidSet},
	}, nil).Wait()
	if err != nil {
		log.Println("Error searching for new messages:", err)
		return
	}
	uids := s.AllUIDs()
	if len(uids) == 0 {
		return
	}

	// Fetch new messages
	fmt.Println("Fetching new messages...", len(uids), uidSet)
	messages, err := mc.client.Fetch(uidSet, &imap.FetchOptions{
		BodySection: []*imap.FetchItemBodySection{{
			Peek: true,
		}},
		Envelope: true,
	}).Collect()
	if err != nil {
		log.Println("Error fetching messages:", err)
		return
	}

	for _, msg := range messages {
		var body []byte
		for sec, val := range msg.BodySection {
			if sec.Specifier == imap.PartSpecifierNone {
				body = val
			}
		}
		mr, err := mail.CreateReader(bytes.NewReader(body))
		if err != nil {
			log.Println("Error creating mail reader:", err)
			continue
		}

		m := &mailstream.Mail{
			UID:     uint32(msg.UID),
			From:    msg.Envelope.From,
			To:      msg.Envelope.To,
			Subject: msg.Envelope.Subject,
			Date:    msg.Envelope.Date,
		}

		for {
			p, err := mr.NextPart()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				log.Println("Error reading mail part:", err)
				break
			}
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				contentType, _, _ := h.ContentType()
				switch contentType {
				case "text/plain":
					m.Plain, _ = io.ReadAll(p.Body)
				case "text/html":
					m.HTML, _ = io.ReadAll(p.Body)
				}
			case *mail.AttachmentHeader:
			}
		}

		// Update lastUID
		if msg.UID > mc.lastUID {
			mc.lastUID = msg.UID
		}

		// Publish the new mail
		mc.Relay.Broadcast(m)
	}
}
