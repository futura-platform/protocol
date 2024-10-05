package basicgroupsprotocol

import (
	"context"
)

type GlobalProvider interface {
	// Deprecated: Use lesser scoped provider methods instead to ensure type safety
	FetchGroup(ctx context.Context, groupType, groupID string) ([]string, error)
	// Deprecated: Use lesser scoped provider methods instead to ensure type safety
	SetGroup(ctx context.Context, groupType, groupID string, entries []string) error
}

type GroupConfig struct {
	EntryTypeSingular string
	EntryTypePlural   string
	EntryPlaceholder  string

	Icon string
}

type Parsable[T any] interface {
	ParseEntry(string) (T, error)
	SerializeEntry() string

	Equals(T) bool

	GetGroupConfig() GroupConfig
}
