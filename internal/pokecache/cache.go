package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entries: make(map[string]CacheEntry), mu: sync.Mutex{}}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = CacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result := c.entries[key].val
	if result == nil {
		return result, false
	}
	return result, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	fmt.Println(<-t.C)

	go func() {
		for range t.C {
			// fmt.Printf("Tick at: %v\n", t.UTC())
			c = &Cache{entries: make(map[string]CacheEntry), mu: sync.Mutex{}} //clear out the cache
		}
	}()
}
