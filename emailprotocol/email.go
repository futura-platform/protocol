package emailprotocol

import (
	"errors"
	"strconv"
	"strings"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
)

type Email struct {
	ImapEmail    string
	ImapPassword string

	Email    string
	Password string
}

var _ basicgroupsprotocol.Parsable[Email] = Email{}

func (Email) ParseEntry(u string) (Email, error) {
	spl := strings.Split(u, ":")
	switch len(spl) {
	case 2:
		return Email{
			ImapEmail: spl[0],
			Email:     spl[0],

			ImapPassword: spl[1],
			Password:     spl[1],
		}, nil
	case 4:
		return Email{
			ImapEmail:    spl[0],
			ImapPassword: spl[1],

			Email:    spl[2],
			Password: spl[3],
		}, nil
	default:
		return Email{}, errors.New("failed to parse email: got " + strconv.Itoa(len(spl)) + " sections")
	}
}

func (e Email) SerializeEntry() string {
	return e.ImapEmail + ":" + e.ImapPassword + ":" + e.Email + ":" + e.Password
}

func (e Email) Equals(e2 Email) bool {
	return e == e2
}

func (e Email) GetGroupConfig() basicgroupsprotocol.GroupConfig {
	return basicgroupsprotocol.GroupConfig{
		EntryTypeSingular: "Email",
		EntryTypePlural:   "Emails",
		EntryPlaceholder:  "imapEmail:imapPassword:email:password or bothEmail:bothPassword",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
  <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
</svg>`,
	}
}
