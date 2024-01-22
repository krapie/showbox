package utils

import lru "github.com/hashicorp/golang-lru"

type LRUCache struct {
	cache *lru.Cache
}

func New() (*LRUCache, error) {
	cache, err := lru.New(128)
	if err != nil {
		return nil, err
	}

	return &LRUCache{
		cache: cache,
	}, nil
}
