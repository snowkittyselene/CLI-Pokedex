package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Mu      *sync.Mutex
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Mu:      &sync.Mutex{},
		Entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	val, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.Mu.Lock()
		for key, val := range c.Entries {
			if t.Compare(val.createdAt.Add(interval)) == 1 {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()
	}
}
