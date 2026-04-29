package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{cache: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	} 
	return value.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for t := range ticker.C {
		c.mu.Lock()
		for key, value := range c.cache {
			if value.createdAt.Before(t.Add(-interval))  {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}


