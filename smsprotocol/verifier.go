package smsprotocol

import (
	"context"

	"github.com/nyaruka/phonenumbers"
)

type Provider interface {
	CreateSMSVerifier(ctx context.Context, providerType string) (Verifier, error)
}

type Verifier interface {
	GetNumber(ctx context.Context, service, country string) (Number, error)
}

type Number interface {
	Phone() *phonenumbers.PhoneNumber

	GetNextCode(context.Context) (string, error)
}
