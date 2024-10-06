//go:build !frontend
// +build !frontend

package emailprotocol

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

type refCountedImapClient struct {
	*client.Client

	e *Email

	refs int
}

func (r *refCountedImapClient) Close() error {
	mutex.Lock()
	defer mutex.Unlock()

	r.refs--
	if r.refs == 0 {
		delete(imapClientCache, r.e.ImapEmail)
		return r.Client.Logout()
	}
	return nil
}

var (
	imapClientCache = make(map[string]*refCountedImapClient)
	mutex           = &sync.Mutex{}
)

func (e *Email) getImapClient() (*refCountedImapClient, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if c, ok := imapClientCache[e.ImapEmail]; ok {
		c.refs++
		return c, nil
	}

	emailDomain := e.ImapEmail[strings.LastIndex(e.ImapEmail, "@")+1:]
	imapHost, ok := knownImapHosts[emailDomain]
	if !ok {
		return nil, fmt.Errorf("unknown email domain: %s", emailDomain)
	}

	c, err := client.DialTLS(fmt.Sprintf("%s:%d", imapHost.host, imapHost.port), &tls.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %w", err)
	}

	// Login
	if err := c.Login(e.ImapEmail, e.ImapPassword); err != nil {
		return nil, fmt.Errorf("error logging in: %w", err)
	}

	rc := &refCountedImapClient{
		Client: c,
		e:      e,
		refs:   1,
	}
	imapClientCache[e.ImapEmail] = rc
	return rc, nil
}

// getEmailsBySender connects to Gmail's IMAP server, authenticates,
// and retrieves emails by the specified sender.
func (e *Email) GetEmailsBySenderAndTime(sender string, since time.Time) ([]string, error) {
	// Set search criteria
	criteria := imap.NewSearchCriteria()
	criteria.Header.Add("From", sender)
	criteria.SentSince = since

	return e.GetEmailsByCriteriaSince(criteria, since)
}

func (e *Email) GetEmailsByCriteriaSince(criteria *imap.SearchCriteria, since time.Time) ([]string, error) {
	c, err := e.getImapClient()
	if err != nil {
		return nil, fmt.Errorf("error getting IMAP client: %w", err)
	}
	defer c.Close()

	// Select INBOX
	_, err = c.Select("INBOX", false)
	if err != nil {
		return nil, fmt.Errorf("error selecting INBOX: %w", err)
	}

	// Perform search
	uids, err := c.Search(criteria)
	if err != nil {
		return nil, fmt.Errorf("error searching emails: %w", err)
	}

	// Fetch specific items
	seqset := new(imap.SeqSet)
	seqset.AddNum(uids...)
	items := []imap.FetchItem{"BODY.PEEK[]", imap.FetchInternalDate}
	messages := make(chan *imap.Message, 10) // buffered channel
	go func() {
		if err := c.Fetch(seqset, items, messages); err != nil {
			log.Println("Fetch error:", err)
		}
	}()

	// Collect email HTML bodies
	var bodies []string
	for msg := range messages {
		if msg.InternalDate.Before(since) {
			continue
		}

		if body := msg.GetBody(&imap.BodySectionName{}); body != nil {
			entity, err := mail.CreateReader(body)
			if err != nil {
				log.Println("Read message body error:", err)
				continue
			}
			for {
				p, err := entity.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Println("Reading part error:", err)
					break
				}

				if strings.Contains(p.Header.Get("Content-Type"), "text/html;") {
					b, err := io.ReadAll(p.Body)
					if err != nil {
						log.Println("Read HTML body error:", err)
						continue
					}
					bodies = append(bodies, string(b))
				}
			}
		} else {
			log.Println("Body is nil")
		}
	}

	return bodies, nil
}
