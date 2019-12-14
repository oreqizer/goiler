package cache

import (
	"context"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	client *cache.Cache
	ttl    time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{client: cache.New(ttl, ttl*2), ttl: ttl}
}

func (c *Cache) Add(ctx context.Context, hash string, query string) {
	c.client.Set(hash, query, cache.DefaultExpiration)
}

func (c *Cache) Get(ctx context.Context, hash string) (string, bool) {
	s, found := c.client.Get(hash)
	if !found {
		return "", false
	}

	return s.(string), true
}
