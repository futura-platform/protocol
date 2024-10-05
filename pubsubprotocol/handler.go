package pubsubprotocol

import (
	"context"

	"github.com/zerofox-oss/go-msg"
)

type ReadMessage struct {
	Attributes msg.Attributes
	Body       []byte
}

type Scope string

const (
	// Send messages to all subscribers within this automation
	ScopeAutomation Scope = "automation"
	// Send messages to all subscribers within this user
	ScopeUser Scope = "local"
)

type Provider interface {
	Publish(context.Context, Scope, msg.Message) error
	Subscribe(ctx context.Context, filter func(msg.Attributes) bool) (chan ReadMessage, chan error)
}
