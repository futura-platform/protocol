package captchaprotocol

import "github.com/futura-platform/protocol/netprotocol"

type DatadomeParams struct {
	Client netprotocol.BaseHttpClient

	WebsiteURL, Referrer string
	CaptchaUrl           string

	UserAgent string
}
