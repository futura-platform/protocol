package captchaprotocol

import (
	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type GeetestParams struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL                string `json:"websiteURL"`
	GT                        string `json:"gt"`
	Challenge                 string `json:"challenge"`
	CaptchaId                 string `json:"captchaId,omitempty"` // for v4
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
}
