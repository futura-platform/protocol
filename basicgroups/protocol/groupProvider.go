package basicgroupsprotocol

import "context"

type Provider[T Parsable[T]] interface {
	// Returns an entry specific to the current task, that won't change between calls
	GetStaticEntry(ctx context.Context) (T, error)
	// Atomically loads entries on a global increment. 1st call will return the first entry, 2nd call will return the 2nd entry, etc.
	LoadIncrementalEntry(ctx context.Context) (T, error)

	ReplaceEntry(ctx context.Context, old T, new T) error
	CheckEntryExistence(ctx context.Context, entry T) (bool, error)

	GetEntries(ctx context.Context) ([]T, error)
}
