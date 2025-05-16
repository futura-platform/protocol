package captchaprotocol

import (
	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
)

type GeetestParamsV3 struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL                string `json:"websiteURL"`
	GT                        string `json:"gt"`
	Challenge                 string `json:"challenge"`
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
}

type GeetestSolutionV3 struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
	Seccode   string `json:"seccode"`
}

type GeetestParamsV4 struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL                string `json:"websiteURL"`
	GT                        string `json:"gt"`
	Challenge                 string `json:"challenge"`
	CaptchaId                 string `json:"captchaId"`
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
}

type GeetestSolutionV4 struct {
	CaptchaId     string `json:"captcha_id"`
	CaptchaOutput string `json:"captcha_output"`
	GenTime       string `json:"gen_time"`
	LotNumber     string `json:"lot_number"`
	PassToken     string `json:"pass_token"`
}
