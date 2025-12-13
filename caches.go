package aklapi

import (
	"github.com/phuslu/lru"
)

const defCacheSz = 100 // seems reasonable.
var (
	addrCache    = newLRUCache[string, *AddrResponse](defCacheSz)
	rubbishCache = rubbishResultCache{lc: newLRUCache[string, *CollectionDayDetailResult](defCacheSz)}
)

type lruCache[K comparable, V any] struct {
	cache *lru.LRUCache[K, V]
}

func newLRUCache[K comparable, V any](size int) *lruCache[K, V] {
	return &lruCache[K, V]{
		cache: lru.NewLRUCache[K, V](size),
	}
}

func (c *lruCache[K, V]) Lookup(key K) (resp V, ok bool) {
	var nothing V
	if NoCache {
		return nothing, false
	}
	return c.cache.Get(key)
}

func (c *lruCache[K, V]) Add(key K, value V) {
	c.cache.Set(key, value)
}

func (c *lruCache[K, V]) Delete(key K) {
	c.cache.Delete(key)
}

type rubbishResultCache struct {
	lc *lruCache[string, *CollectionDayDetailResult]
}

func (c *rubbishResultCache) Lookup(searchText string) (result *CollectionDayDetailResult, ok bool) {
	result, ok = c.lc.Lookup(searchText)
	if !ok {
		return nil, false
	}

	today := now()
	for _, res := range result.Collections {
		if today.After(res.Date) || res.Date.IsZero() {
			// invalidate from cache.
			c.lc.Delete(searchText)
			return nil, false
		}
	}
	return
}

func (c *rubbishResultCache) Add(searchText string, result *CollectionDayDetailResult) {
	c.lc.Add(searchText, result)
}
