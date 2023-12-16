package repository

import (
	"e-wallet/domain"
	"e-wallet/internal/config"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

type redisCacheRepository struct {
	rdb *redis.Client
}

func NewRedisClient(cnf *config.Config) domain.CacheRepository {
	return &redisCacheRepository{
		rdb: redis.NewClient(&redis.Options{
			Addr:     cnf.Redis.Addr,
			Password: cnf.Redis.Password,
			DB:       0,
		}),
	}
}

func (r redisCacheRepository) Get(key string) ([]byte, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r redisCacheRepository) Set(key string, entry []byte) error {
	return r.rdb.Set(context.Background(), key, entry, 15*time.Minute).Err()
}
