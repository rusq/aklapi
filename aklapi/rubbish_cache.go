package aklapi

import "time"

type rubbishResultCache map[string]*CollectionDayDetailResult

var rubbishCache rubbishResultCache = make(rubbishResultCache, 0)

var now = time.Now // time generator

func (c rubbishResultCache) Lookup(searchText string) (result *CollectionDayDetailResult, ok bool) {
	result, ok = c[searchText]
	if !ok {
		return
	}

	today := now()
	for _, res := range result.Collections {
		if today.After(res.Date) || res.Date.IsZero() {
			// invalidate from cache.
			delete(c, searchText)
			return nil, false
		}
	}
	return
}

func (c rubbishResultCache) Add(searchText string, result *CollectionDayDetailResult) {
	c[searchText] = result
}
