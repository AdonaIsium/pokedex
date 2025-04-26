package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	Mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CacheMap: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.Mu.Lock()
	cache.CacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
	cache.Mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	if val, exists := cache.CacheMap[key]; exists {
		return val.val, true
	} else {
		return nil, false
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		cutoff := time.Now().Add(interval)

		cache.Mu.Lock()
		for k, item := range cache.CacheMap {
			if item.createdAt.Before(cutoff) {
				delete(cache.CacheMap, k)
			}
		}
		cache.Mu.Unlock()
	}
}
