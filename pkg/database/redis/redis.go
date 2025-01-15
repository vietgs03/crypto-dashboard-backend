package redis

import (
	"context"
	"errors"
	"fmt"
	redisV9 "github.com/redis/go-redis/v9"
	"log/slog"
)

var ctx = context.Background()

type redis struct {
	cfg    *ConfigRedis
	client *redisV9.Client
}

var _ EngineRedis = (*redis)(nil)

func NewRedisClient(cfg *ConfigRedis) (EngineRedis, error) {
	redis := &redis{
		cfg: cfg,
	}
	urlRedis := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
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

func (r *redis) Get(key string) ([]byte, bool, error) {
	byteValue, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redisV9.Nil) {
		return nil, false, err
	}
	if err != nil {
		return nil, false, err
	}
	return byteValue, true, nil
}
