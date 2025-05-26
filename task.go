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
	"github.com/futura-platform/protocol/netprotocol/proxyprotocol"
	"github.com/futura-platform/protocol/pubsubprotocol"
	"github.com/futura-platform/protocol/sessionsprotocol"
	"github.com/futura-platform/protocol/settingsprotocol"
	"github.com/futura-platform/protocol/smsprotocol"
	"github.com/futura-platform/protocol/userinputprotocol"
)

type BaseTask interface {
	flowprotocol.Context
	logprotocol.Logger
	netprotocol.BaseHttpClient
	captchaprotocol.Provider
	browserprotocol.Spawner
	pubsubprotocol.Provider
	sessionsprotocol.Provider
	settingsprotocol.Provider
	smsprotocol.Provider
	userinputprotocol.Provider
	// basicgroupsprotocol.GenericProvider
	ProxyProvider() basicgroupsprotocol.Provider[*proxyprotocol.Proxy]

	// extendable error delay
	GetErrorDelay() time.Duration

	// step lifecycle hooks. The pointer patameters of these methods are gauranteed to be non-nil, and are edittable.
	BeforeStep(step *flowprotocol.TaskStep) error
	AfterStep(step flowprotocol.TaskStep, result *flowprotocol.TaskStepResult) error

	// wraps goroutine spawning so that recovery is handled and properly logged
	Go(func())

	// step flow
	GetSteps() []flowprotocol.TaskStep

	CurrentStepIndex() int
	CurrentStep() flowprotocol.TaskStep

	// logging
	BLog() *log.Logger
	Fatalf(format string, args ...any) error

	ReturnBasicStepSuccess() flowprotocol.TaskStepResult
	ReturnSmallErrorf(format string, args ...any) flowprotocol.TaskStepResult
	ReturnFatalErrorf(format string, args ...any) flowprotocol.TaskStepResult

	// add a custom column to the frontend task table with the value. Rows without this set will have an empty string in the column.
	SetFrontendColumn(sortKey int, columnName, value string)

	// other
	WithContext(ctx context.Context) BaseTask

	RotateProxy() error
}

// this is the type that users of the protocol package should use
type Task[T any] struct {
	BaseTask

	Params *T
}
