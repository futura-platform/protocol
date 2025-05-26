package flowprotocol

import (
	"context"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Context interface {
	context.Context

	TaskId() string
	FatalError(error)

	Sleep(time.Duration)
}

type TaskStepResult struct {
	// label of the next TaskStep to execute (or terminate to stop the task), nil goes to next step in the flow if success
	NextStepLabel string
	Err           error
}

type TaskStep struct {
	// optional, will be inferred from the runtime function name if not provided
	Label    string
	StepFunc func() TaskStepResult

	// optional, all previous steps that dont have a grouping will be grouped together until another step with a grouping is reached
	StepGrouping string
}

func (step TaskStep) GetStepLabel() string {
	if step.Label == "" {
		funcName := runtime.FuncForPC(reflect.ValueOf(step.StepFunc).Pointer()).Name()
		spl := strings.Split(funcName, ".")
		if len(spl) == 0 {
			return ""
		}
		methodSegment := spl[len(spl)-1]
		spl2 := strings.Split(methodSegment, "-")
		if len(spl2) == 0 {
			return ""
		}
		step.Label = spl2[0]
	}

	return step.Label
}
