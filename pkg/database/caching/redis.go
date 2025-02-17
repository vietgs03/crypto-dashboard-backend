package caching

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"crypto-dashboard/pkg/settings"

	redisV9 "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type CacheClient struct {
	cfg    *settings.CacheSetting
	client *redisV9.Client
}

func NewRedisClient(cfg *settings.CacheSetting) (*CacheClient, error) {
	redis := &CacheClient{
		cfg: cfg,
	}
	urlRedis := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	redis.client = redisV9.NewClient(&redisV9.Options{
		Addr:     urlRedis,
		Password: cfg.Password,
		DB:       cfg.Database,
		PoolSize: cfg.PoolSize,
	})

	_, err := redis.client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("failed to connect redis" + err.Error())
	}
	slog.Info("redis connect success")
	return redis, nil
}
