package cache

import (
	"context"
	"github.com/patrickmn/go-cache"
	"time"
)

// Cache holds cache information
type Cache struct {
	client *cache.Cache
	ttl    time.Duration
}

// New creates a new cache
func New(ttl time.Duration) *Cache {
	return &Cache{client: cache.New(ttl, ttl*2), ttl: ttl}
}

// Add adds an entry to cache
func (c *Cache) Add(ctx context.Context, hash string, query string) {
	c.client.Set(hash, query, cache.DefaultExpiration)
}

// Get gets an entry from cache
func (c *Cache) Get(ctx context.Context, hash string) (string, bool) {
	s, found := c.client.Get(hash)
	if !found {
		return "", false
	}

	return s.(string), true
}
