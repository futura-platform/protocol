package protocol

import (
	"context"
	"log"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
	"github.com/futura-platform/protocol/browserprotocol"
	"github.com/futura-platform/protocol/captchaprotocol"
	"github.com/futura-platform/protocol/flowprotocol"
	"github.com/futura-platform/protocol/httpsserveprotocol"
	"github.com/futura-platform/protocol/logprotocol"
	"github.com/futura-platform/protocol/netprotocol"
	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
	"github.com/futura-platform/protocol/pubsubprotocol"
	"github.com/futura-platform/protocol/sessionsprotocol"
	"github.com/futura-platform/protocol/settingsprotocol"
	"github.com/futura-platform/protocol/smsprotocol"
	"github.com/futura-platform/protocol/userinputprotocol"
)

type BaseTask interface {
	// step flow
	flowprotocol.Context
	// logging
	logprotocol.Logger
	// basic logging
	BLog() *log.Logger

	netprotocol.BaseHttpClient
	captchaprotocol.Provider
	browserprotocol.Spawner
	pubsubprotocol.Provider
	sessionsprotocol.Provider
	settingsprotocol.Provider
	smsprotocol.Provider
	userinputprotocol.Provider
	flowprotocol.LifecycleHooks
	httpsserveprotocol.Provider

	// basicgroupsprotocol.GenericProvider
	ProxyProvider() basicgroupsprotocol.Provider[*proxyprotocol.Proxy]
	RotateProxy() error

	// other
	WithContext(ctx context.Context) BaseTask
}

// this is the type that users of the protocol package should use
type Task[T any] struct {
	BaseTask
	netprotocol.HttpClient

	Params *T
}
