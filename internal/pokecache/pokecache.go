package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val,
		time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cachE, ok := c.cache[key]
	return cachE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	beforeTime := time.Now().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(beforeTime) {
			delete(c.cache, k)
		}
	}
}
