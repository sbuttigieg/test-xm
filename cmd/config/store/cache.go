package store

import (
	"github.com/go-redis/redis"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/store"
	"github.com/sbuttigieg/test-xm/xm_app/store/cache"
)

// NewInmem create new in memory store
func NewCache(cfg *app.Config, redis redis.UniversalClient) store.Cache {
	return cache.New(cfg, redis)
}
