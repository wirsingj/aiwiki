package cache

import (
	"sync"
	"time"
)

type item[T any] struct {
	value     T
	expiresAt time.Time
}

type Cache[T any] struct {
	mu    sync.RWMutex
	ttl   time.Duration
	items map[string]item[T]
}

func New[T any](ttl time.Duration) *Cache[T] {
	return &Cache[T]{
		ttl:   ttl,
		items: make(map[string]item[T]),
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	entry, ok := c.items[key]
	c.mu.RUnlock()

	var zero T
	if !ok {
		return zero, false
	}
	if time.Now().After(entry.expiresAt) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return zero, false
	}
	return entry.value, true
}

func (c *Cache[T]) Set(key string, value T) {
	c.mu.Lock()
	c.items[key] = item[T]{
		value:     value,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()
}
