package captchaprotocol

import "github.com/futura-platform/protocol/netprotocol/proxyprotocol"

type AWSWAFTokenParams struct {
	Proxy *proxyprotocol.Proxy `json:"-"`

	WebsiteURL     string `json:"websiteURL"`
	AWSKey         string `json:"awsKey,omitempty"`
	AWSIv          string `json:"awsIv,omitempty"`
	AWSContext     string `json:"awsContext,omitempty"`
	AWSChallengeJS string `json:"awsChallengeJS,omitempty"`
}
