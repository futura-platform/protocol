package captchaprotocol

import "github.com/futura-platform/protocol/netprotocol"

type TurnstileParams struct {
	Proxy *netprotocol.Proxy

	WebsiteURL string
	WebsiteKey string
	UserAgent  string

	Action   string
	CData    string
	PageData string
}
