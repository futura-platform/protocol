package logprotocol

type Logger interface {
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

	// add a custom column to the frontend task table with the value. Rows without this set will have an empty string in the column.
	SetFrontendColumn(sortKey int, columnName, value string)
}
