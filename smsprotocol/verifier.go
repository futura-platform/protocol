package smsprotocol

import (
	"context"

	"github.com/nyaruka/phonenumbers"
)

type Provider interface {
	CreateSMSVerifier(ctx context.Context, providerType string) (Verifier, error)
}

type Verifier interface {
	GetCode(ctx context.Context, service, country string, onPhoneNumber func(phoneNumber *phonenumbers.PhoneNumber) error) (code string, err error)
}
