package userinputprotocol

import (
	"context"

	"github.com/futura-platform/protocol/flowprotocol"
)

type InputField struct {
	Key string

	DefaultValue *string

	Label           string
	ValidationRegex string
}

type Provider interface {
	GetUserFormInput(ctx flowprotocol.Context, typeId, imageUrl string, details map[string]string, inputs []InputField) (UserFormInput, error)
}

type UserFormInput interface {
	Close()
	ForEachInput(f func(key string, value string))
	GetInput(key string) string
	InputLen() int
	WaitForInput(ctx context.Context)
}
