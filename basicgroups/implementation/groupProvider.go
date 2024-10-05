package basicgroupsimplementation

import (
	"context"
	"fmt"
	"sync/atomic"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
)

type cachedBasicGroupProvider[T basicgroupsprotocol.Parsable[T]] struct {
	groupsProvider basicgroupsprotocol.GroupsProvider[T]

	accessIncrement atomic.Int64
	groupId         string
}

// Deprecated: Define a provider on a parameter type to ensure type safety. This function is just a helper for auto generated code.
func NewCachedBasicGroupProvider[T basicgroupsprotocol.Parsable[T]](
	groupsProvider basicgroupsprotocol.GroupsProvider[T],
	groupId string,
	startOffset int64,
) basicgroupsprotocol.Provider[T] {
	r := &cachedBasicGroupProvider[T]{
		groupsProvider: groupsProvider,
		groupId:        groupId,
	}
	r.accessIncrement.Store(startOffset)

	return r
}

func (p *cachedBasicGroupProvider[T]) getThisGroup(ctx context.Context) ([]T, error) {
	return p.groupsProvider.GetGroup(ctx, p.groupId)
}

func (p *cachedBasicGroupProvider[T]) LoadEntry(ctx context.Context) (zeroVal T, err error) {
	entries, err := p.getThisGroup(ctx)
	if err != nil {
		return zeroVal, err
	} else if len(entries) == 0 {
		return zeroVal, nil
	}

	index := (p.accessIncrement.Add(1) - 1) % int64(len(entries))
	return entries[index], nil
}

func (p *cachedBasicGroupProvider[T]) ReplaceEntry(ctx context.Context, old T, new T) error {
	entries, err := p.getThisGroup(ctx)
	if err != nil {
		return err
	}

	for i, e := range entries {
		if e.Equals(old) {
			entries[i] = new
			return p.groupsProvider.SetGroup(ctx, p.groupId, entries)
		}
	}

	return fmt.Errorf("entry not found")
}

func (p *cachedBasicGroupProvider[T]) CheckEntryExistence(ctx context.Context, entry T) (bool, error) {
	entries, err := p.getThisGroup(ctx)
	if err != nil {
		return false, err
	}

	for _, e := range entries {
		if e.Equals(entry) {
			return true, nil
		}
	}

	return false, nil
}
