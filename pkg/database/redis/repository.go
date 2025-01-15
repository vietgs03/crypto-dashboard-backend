package redis

type EngineRedis interface {
	Get(key string) ([]byte, bool, error)
}
