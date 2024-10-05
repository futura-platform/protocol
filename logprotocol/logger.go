package logprotocol

type Logger interface {
	SetStep(step int)

	LargeSuccess(args ...any)
	LargeSuccessf(format string, args ...any)

	SmallSuccess(args ...any)
	SmallSuccessf(format string, args ...any)

	Info(args ...any)
	Infof(format string, args ...any)

	SmallError(args ...any)
	SmallErrorf(format string, args ...any)

	LargeError(args ...any)
	LargeErrorf(format string, args ...any)

	SetFinal()
}
