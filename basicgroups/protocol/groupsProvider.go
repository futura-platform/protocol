package basicgroupsprotocol

import "context"

type GroupsProvider[T Parsable[T]] interface {
	GetGroup(ctx context.Context, groupID string) ([]T, error)
	SetGroup(ctx context.Context, groupID string, entries []T) error
}
