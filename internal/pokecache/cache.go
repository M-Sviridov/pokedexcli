package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  map[string]cacheEntry{},
		mu:       sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	entry, ok := cache.entries[key]
	cache.mu.Unlock()
	return entry.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for key, entry := range cache.entries {
			age := time.Since(entry.createdAt)
			if age > cache.interval {
				cache.mu.Lock()
				delete(cache.entries, key)
				cache.mu.Unlock()
			}
		}
	}
}
