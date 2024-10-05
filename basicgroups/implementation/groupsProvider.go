package basicgroupsimplementation

import (
	"context"
	"fmt"
	"time"

	basicgroupsprotocol "github.com/futura-platform/protocol/basicgroups/protocol"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/sync/singleflight"
)

type cachedBasicGroupsProvider[T basicgroupsprotocol.Parsable[T]] struct {
	g      basicgroupsprotocol.GlobalProvider
	typeId string

	cacheQueries singleflight.Group
	cache        *xsync.MapOf[string, cachedGroup[T]]

	maxAge time.Duration
}

// Deprecated: Define a provider on a parameter type to ensure type safety. This function is just a helper for auto generated code.
func NewCachedBasicGroupsProvider[T basicgroupsprotocol.Parsable[T]](
	genericProvider basicgroupsprotocol.GlobalProvider, typeId string, maxAge time.Duration,
) basicgroupsprotocol.GroupsProvider[T] {
	return &cachedBasicGroupsProvider[T]{
		g:      genericProvider,
		typeId: typeId,

		cache:  xsync.NewMapOf[string, cachedGroup[T]](),
		maxAge: maxAge,
	}
}

type cachedGroup[T basicgroupsprotocol.Parsable[T]] struct {
	entries     []T
	lastUpdated time.Time
}

func (p *cachedBasicGroupsProvider[T]) GetGroup(ctx context.Context, groupID string) ([]T, error) {
	result, err, _ := p.cacheQueries.Do(groupID, func() (any, error) {
		if cached, ok := p.cache.Load(groupID); ok && time.Since(cached.lastUpdated) < p.maxAge {
			return cached.entries, nil
		}

		rawEntries, err := p.g.FetchGroup(ctx, p.typeId, groupID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch group: %w", err)
		}

		var parser T
		entries := make([]T, len(rawEntries))
		for i, rawEntry := range rawEntries {
			entry, err := parser.ParseEntry(rawEntry)
			if err != nil {
				return nil, fmt.Errorf("failed to parse entry [%d]: %w", i, err)
			}
			entries[i] = entry
		}

		p.cache.Store(groupID, cachedGroup[T]{entries: entries, lastUpdated: time.Now()})
		return entries, nil
	})
	if err != nil {
		return nil, err
	}

	return result.([]T), nil
}

func (p *cachedBasicGroupsProvider[T]) SetGroup(ctx context.Context, groupID string, entries []T) error {
	rawEntries := make([]string, len(entries))
	for i, entry := range entries {
		rawEntries[i] = entry.SerializeEntry()
	}
	err := p.g.SetGroup(ctx, p.typeId, groupID, rawEntries)
	if err != nil {
		return fmt.Errorf("failed to set group: %w", err)
	}

	p.cache.Store(groupID, cachedGroup[T]{entries: entries, lastUpdated: time.Now()})
	return nil
}
