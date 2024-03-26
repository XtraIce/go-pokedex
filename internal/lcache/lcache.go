package lcache

import (
	"sync"
	"time"
)

// Cache represents a local cache.
type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
}

// cacheEntry represents an entry in the cache.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache creates a new Cache instance with the specified reaping interval.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval) // start reaping goroutine
	return cache
}

// Add adds a new key-value pair to the cache.
func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves the value associated with the specified key from the cache.
// It returns the value and a boolean indicating whether the key was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// reapLoop is a goroutine that periodically reaps expired cache entries.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.reap(interval)
	}
}

// reap removes expired cache entries from the cache.
func (c *Cache) reap(interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.cache {
		if time.Since(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
