package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

const companies string = "companies"

type cacheMiddleware struct {
	cache  store.Cache
	config *app.Config
	next   services.Store
}

// NewCacheMiddleware creates a new cache middleware.
func NewCacheMiddleware(cfg *app.Config, next services.Store, cache store.Cache) services.Store {
	m := cacheMiddleware{
		cache:  cache,
		config: cfg,
		next:   next,
	}

	return &m
}

func (m *cacheMiddleware) Create(ctx context.Context, req *models.Company) (string, error) {
	resp, err := m.next.Create(ctx, req)
	if err != nil {
		return "", errors.Wrapf(err, "")
	}

	return resp, nil
}

func (m *cacheMiddleware) Get(ctx context.Context, id string) (*models.Company, error) {
	key := fmt.Sprintf("%s_%s", id, companies)

	p, ok := m.cache.GetKeyBytes(key)
	if ok {
		var cacheValue *models.Company

		err := json.Unmarshal(p, &cacheValue)
		if err != nil {
			return nil, err
		}

		return cacheValue, nil
	}

	company, err := m.next.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	err = m.cache.SetKey(key, company, m.config.CacheExpiry)
	if err != nil {
		return nil, err
	}

	return company, nil
}
