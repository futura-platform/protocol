package flowprotocol

import "context"

type UserInput interface {
	WaitForInput(context.Context)

	GetInput(key string) string
	InputLen() int
	ForEachInput(func(key, value string))

	Close()
}
