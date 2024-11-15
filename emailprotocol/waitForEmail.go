//go:build !frontend
// +build !frontend

package emailprotocol

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/robin-samuel/mailstream"
)

type imapHost struct {
	host string
	port int
}

type mailStreamRef struct {
	client              *MailClient
	cancelClientContext context.CancelFunc
	refs                int
	errs                <-chan error
}

var (
	activeEmailStreams      = make(map[string]*mailStreamRef)
	activeEmailStreamsMutex sync.Mutex

	knownImapHosts = map[string]imapHost{
		"gmail.com": {
			host: "imap.gmail.com",
			port: 993,
		},
	}
)

type EmailStream interface {
	Subscribe() <-chan *mailstream.Mail
}

func getEmailStream(ctx context.Context, imapEmail, imapPassword string) (stream *MailClient, errs <-chan error, err error) {
	activeEmailStreamsMutex.Lock()
	defer activeEmailStreamsMutex.Unlock()
	defer func() {
		go func() {
			<-ctx.Done()
			activeEmailStreamsMutex.Lock()
			defer activeEmailStreamsMutex.Unlock()

			if stream, ok := activeEmailStreams[imapEmail]; ok {
				stream.refs--
				if stream.refs == 0 {
					delete(activeEmailStreams, imapEmail)
					stream.client.Close()
					stream.cancelClientContext()
				}
			}
		}()
	}()

	if stream, ok := activeEmailStreams[imapEmail]; ok {
		stream.refs++
		return stream.client, stream.errs, nil
	}

	emailDomain := imapEmail[strings.LastIndex(imapEmail, "@")+1:]
	imapHost, ok := knownImapHosts[emailDomain]
	if !ok {
		return nil, nil, fmt.Errorf("unknown email domain: %s", imapEmail)
	}

	updateCtx, cancel := context.WithCancel(context.Background())
	stream, err = newMailClient(updateCtx, imapHost.host, imapHost.port, imapEmail, imapPassword)
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error creating email stream: %w", err)
	}
	activeEmailStreams[imapEmail] = &mailStreamRef{
		client:              stream,
		cancelClientContext: cancel,
		refs:                1,
		errs:                errs,
	}

	return stream, errs, nil
}

func (e *Email) WaitForEmail(ctx context.Context, filter func(*mailstream.Mail) bool) (*mailstream.Mail, error) {
	client, errChan, err := getEmailStream(ctx, e.ImapEmail, e.ImapPassword)
	if err != nil {
		return nil, fmt.Errorf("error getting email stream: %w", err)
	}

	listener := client.Listener(1)
	defer listener.Close()

	for {
		select {
		case err := <-errChan:
			return nil, fmt.Errorf("error in email stream: %w", err)
		case <-ctx.Done():
			return nil, ctx.Err()
		case mail, ok := <-listener.Ch():
			if !ok {
				return nil, fmt.Errorf("email stream closed")
			} else if filter(mail) {
				return mail, nil
			}
		}
	}
}
