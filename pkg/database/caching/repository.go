package caching

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

type EngineCaching interface {
	Get(key string) ([]byte, bool, error)
}

func (r *CacheClient) Get(key string) ([]byte, bool, error) {
	byteValue, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, false, err
	}
	if err != nil {
		return nil, false, err
	}
	return byteValue, true, nil
}
