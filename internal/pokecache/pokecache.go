package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	savedEntries map[string]cacheEntry
	pkmu         *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		savedEntries: make(map[string]cacheEntry),
		pkmu:         &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c Cache) Add(key string, val []byte) {
	c.pkmu.Lock()
	defer c.pkmu.Unlock()

	c.savedEntries[key] = cacheEntry{time.Now(), val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.pkmu.Lock()
	defer c.pkmu.Unlock()

	entry, exists := c.savedEntries[key]

	return entry.val, exists
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.pkmu.Lock()
	defer c.pkmu.Unlock()
	for k, v := range c.savedEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.savedEntries, k)
		}
	}
}
