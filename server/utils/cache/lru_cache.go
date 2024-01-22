package cache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cache struct {
	cache *expirable.LRU[any, any]
}

func New() *Cache {
	cache := expirable.NewLRU[any, any](1000, nil, time.Minute*5)

	return &Cache{
		cache: cache,
	}
}

func (c *Cache) Add(key any, value any) {
	c.cache.Add(key, value)
}

func (c *Cache) Get(key any) (any, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Remove(key any) {
	c.cache.Remove(key)
}

func (c *Cache) Exists(id primitive.ObjectID) bool {
	return c.cache.Contains(id)
}
