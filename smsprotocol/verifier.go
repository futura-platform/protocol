package smsprotocol

import (
	"context"

	"github.com/nyaruka/phonenumbers"
)

// Deprecated: for internal use only, define a Verifier to your params struct
// and let it be populated automatically
type Provider interface {
	CreateSMSVerifier(ctx context.Context, providerType string) (Verifier, error)
}

type Verifier interface {
	GetCode(ctx context.Context, service, country string, onPhoneNumber func(phoneNumber *phonenumbers.PhoneNumber) error) (code string, err error)
}
