package cache

import (
	"github.com/go-redis/redis"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

// New create new redis store
func New(c *app.Config, db redis.UniversalClient) store.Cache {
	s := &cache{
		config: c,
		db:     db,
	}

	return s
}

type cache struct {
	config *app.Config
	db     redis.UniversalClient
}
