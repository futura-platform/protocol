package captchaprotocol

import "github.com/futura-platform/protocol/netprotocol/proxyprotocol"

type TurnstileParams struct {
	Proxy *proxyprotocol.Proxy

	WebsiteURL string
	WebsiteKey string
	UserAgent  string

	Action   string
	CData    string
	PageData string
}
