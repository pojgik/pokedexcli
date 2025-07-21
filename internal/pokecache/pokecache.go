package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
} // cacheEntry

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
} // Cache

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: map[string]cacheEntry{},
		mu:      &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
} // NewCache

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
} // Add

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	} else {
		return nil, false
	} // if
} // Get

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.mu.Lock()
		for key := range c.entries {
			if interval < t.Sub(c.entries[key].createdAt) {
				delete(c.entries, key)
			}
		} // for key
		c.mu.Unlock()
	} // for
} // reapLoop
