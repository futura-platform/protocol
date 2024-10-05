package flowprotocol

import (
	"context"
	"time"
)

type TaskStepResult struct {
	NextStepLabel string // label of the next TaskStep to execute (or terminate to stop the task), nil goes to next step in the flow if success
	Success       bool
	Error         error
}

type TaskStep struct {
	// optional, will be inferred from the runtime function name if not provided
	Label    string
	StepFunc func() TaskStepResult

	// optional, all previous steps that dont have a grouping will be grouped together until another step with a grouping is reached
	StepGrouping string
}

type Context interface {
	context.Context

	TaskId() string
	FatalError(error)

	Sleep(time.Duration)
}
