package captchaprotocol

import (
	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type HcaptchaParams struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL         string `json:"websiteURL"`
	WebsiteKey         string `json:"websiteKey"`
	IsInvisible        bool   `json:"isInvisible,omitempty"`
	Data               string `json:"data,omitempty"`
	UserAgent          string `json:"userAgent,omitempty"`
	Cookies            string `json:"cookies,omitempty"` // Format: name1=value1; name2=value2
	FallbackToActualUA bool   `json:"fallbackToActualUA,omitempty"`
}
