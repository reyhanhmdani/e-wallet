package component

import (
	"e-wallet/domain"
	"github.com/allegro/bigcache/v3"
	"log"
	"time"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Fatalf("error when connect cache %s", err.Error())
	}
	return cache
}
