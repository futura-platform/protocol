package flowprotocol

type LifecycleHooks interface {
	// BeforeStep is called before executing a step. The step parameter is guaranteed to be non-nil and can be modified.
	BeforeStep(step *TaskStep) error
	// AfterStep is called after executing a step. the result parameter is guaranteed to be non-nil and can be modified.
	AfterStep(step TaskStep, result *TaskStepResult) error
}
