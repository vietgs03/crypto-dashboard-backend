package caching_test

import (
	"testing"

	"crypto-dashboard/pkg/database/caching"
	"crypto-dashboard/pkg/settings"

	"github.com/stretchr/testify/require"
)

func connectRedisTest() (caching.EngineCaching, error) {
	cfg := settings.CacheSetting{
		Host:     "xxx.x.x.x",
		Port:     6379,
		Password: "xxxxx",
		Database: 0,
		PoolSize: 0,
	}
	engineRedis, err := caching.NewRedisClient(&cfg)
	if err != nil {
		return nil, err
	}
	return engineRedis, nil
}

func TestNewRedisClient(t *testing.T) {
	engineRedis, err := connectRedisTest()
	require.Error(t, err)
	require.Nil(t, engineRedis)
	require.NotEmpty(t, engineRedis)
}
