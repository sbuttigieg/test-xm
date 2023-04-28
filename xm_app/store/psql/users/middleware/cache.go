package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services/users"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

const usersKey string = "users"

type cacheMiddleware struct {
	cache  store.Cache
	config *app.Config
	next   users.Store
}

// NewCacheMiddleware creates a new cache middleware.
func NewCacheMiddleware(cfg *app.Config, next users.Store, cache store.Cache) users.Store {
	m := cacheMiddleware{
		cache:  cache,
		config: cfg,
		next:   next,
	}

	return &m
}

func (m *cacheMiddleware) Create(ctx context.Context, req *models.User) (string, error) {
	resp, err := m.next.Create(ctx, req)
	if err != nil {
		return "", errors.Wrapf(err, "")
	}

	return resp, nil
}

func (m *cacheMiddleware) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	key := fmt.Sprintf("%s_%s", email, usersKey)

	p, ok := m.cache.GetKeyBytes(key)
	if ok {
		var cacheValue *models.User

		err := json.Unmarshal(p, &cacheValue)
		if err != nil {
			return nil, err
		}

		return cacheValue, nil
	}

	user, err := m.next.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	err = m.cache.SetKey(key, user, m.config.CacheExpiry)
	if err != nil {
		return nil, err
	}

	return user, nil
}
