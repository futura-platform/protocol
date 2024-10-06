package captchaprotocol

import "github.com/futura-platform/protocol/netprotocol/proxyprotocol"

type RecaptchaV2Params struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL          string `json:"websiteURL"`
	WebsiteKey          string `json:"websiteKey"`
	RecaptchaDataSValue string `json:"recaptchaDataSValue,omitempty"`
	UserAgent           string `json:"userAgent,omitempty"`
	Cookies             string `json:"cookies,omitempty"` // Format: name1=value1; name2=value2
	IsInvisible         bool   `json:"isInvisible,omitempty"`
}

type RecaptchaV3Params struct {
	WebsiteURL   string `json:"websiteURL"`
	WebsiteKey   string `json:"websiteKey"`
	MinScore     string `json:"minScore,omitempty"`
	PageAction   string `json:"pageAction,omitempty"`
	IsEnterprise bool   `json:"isEnterprise,omitempty"`
	ApiDomain    string `json:"apiDomain,omitempty"`
}
