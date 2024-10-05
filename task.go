package protocol

import (
	"context"
	"log"
	"time"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
	"github.com/futura-platform/protocol/browserprotocol"
	"github.com/futura-platform/protocol/captchaprotocol"
	"github.com/futura-platform/protocol/flowprotocol"
	"github.com/futura-platform/protocol/logprotocol"
	"github.com/futura-platform/protocol/netprotocol"
	"github.com/futura-platform/protocol/pubsubprotocol"
	"github.com/futura-platform/protocol/sessionsprotocol"
	"github.com/futura-platform/protocol/settingsprotocol"
)

type BaseTask interface {
	flowprotocol.Context
	logprotocol.Logger
	netprotocol.BaseHttpClient
	captchaprotocol.Solver
	browserprotocol.Spawner
	pubsubprotocol.Provider
	sessionsprotocol.Provider
	settingsprotocol.Provider
	// basicgroupsprotocol.GenericProvider
	ProxyProvider() basicgroupsprotocol.Provider[*netprotocol.Proxy]

	// extendable getters
	GetErrorDelay() time.Duration
	HandleConsecutiveFails(errs []error) (bool, string)

	// helpers
	// logging
	BLog() *log.Logger
	Fatalf(format string, args ...any) error

	ReturnBasicStepSuccess() flowprotocol.TaskStepResult
	ReturnSmallErrorf(format string, args ...any) flowprotocol.TaskStepResult
	ReturnFatalErrorf(format string, args ...any) flowprotocol.TaskStepResult

	// other
	WithContext(ctx context.Context) BaseTask

	RotateProxy() error
	SpawnSingleTabBrowser(proxy *netprotocol.Proxy) (*browserprotocol.SingleTabBrowser, context.CancelFunc, error)
}

// this is the type that users of the protocol package should use
type Task[T any] struct {
	BaseTask

	Params *T
}
