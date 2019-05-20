package storage

import (
	"github.com/go-redis/redis"
	"time"
)

type Storage interface {
	Set(key, val string, expiration time.Duration) error
	Get(key string) (val string, err error)
}

type RedisStorage struct {
	redisCli redis.UniversalClient
}

func NewRedisStorage(cli redis.UniversalClient) Storage {
	return &RedisStorage{redisCli: cli}
}

func (s *RedisStorage) Set(key, val string, expiration time.Duration) error {
	return s.redisCli.Set(key, val, expiration).Err()
}

func (s *RedisStorage) Get(key string) (val string, err error) {
	return s.redisCli.Get(key).Result()
}
