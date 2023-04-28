package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

const companiesKey string = "companies"

type cacheMiddleware struct {
	cache  store.Cache
	config *app.Config
	next   companies.Store
}

// NewCacheMiddleware creates a new cache middleware.
func NewCacheMiddleware(cfg *app.Config, next companies.Store, cache store.Cache) companies.Store {
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

func (m *cacheMiddleware) Delete(ctx context.Context, id string) error {
	err := m.next.Delete(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "")
	}

	return nil
}

func (m *cacheMiddleware) Get(ctx context.Context, id string) (*models.Company, error) {
	key := fmt.Sprintf("%s_%s", id, companiesKey)

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

func (m *cacheMiddleware) Update(ctx context.Context, id string, data *models.Company, fields []string) (*models.Company, error) {
	resp, err := m.next.Update(ctx, id, data, fields)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	return resp, nil
}
