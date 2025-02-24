package proxy

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Data      []byte
	ExpiresAt time.Time
}

type Cache struct {
	store sync.Map
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{ttl: ttl}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.store.Load(key)
	if !ok {
		return nil, false
	}

	entry := value.(CacheEntry)
	if time.Now().After(entry.ExpiresAt) {
		c.store.Delete(key)
		return nil, false
	}

	return entry.Data, true
}

func (c *Cache) Set(key string, data []byte) {
	c.store.Store(key, CacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(c.ttl),
	})
}
