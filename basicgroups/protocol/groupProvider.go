package basicgroupsprotocol

import "context"

type Provider[T Parsable[T]] interface {
	LoadEntry(ctx context.Context) (T, error)
	ReplaceEntry(ctx context.Context, old T, new T) error
	CheckEntryExistence(ctx context.Context, entry T) (bool, error)

	GetEntries(ctx context.Context) ([]T, error)
}
