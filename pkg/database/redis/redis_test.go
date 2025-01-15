package redis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func connectRedisTest() (EngineRedis, error) {
	cfg := ConfigRedis{
		Host:     "xxx.x.x.x",
		Port:     "xxxx",
		Password: "xxxxx",
		Database: 0,
		PoolSize: 0,
	}
	engineRedis, err := NewRedisClient(&cfg)
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
