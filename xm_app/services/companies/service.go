package companies

import (
	"context"
	"time"

	"github.com/google/uuid"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service
type Service interface {
	Create(context.Context, *models.Company) (string, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (*models.Company, error)
	Update(context.Context, string, *models.Company, []string) (*models.Company, error)
}

func New(config *app.Config, cache store.Cache, store Store, uuidFunc func() uuid.UUID, timeFunc func() time.Time) Service {
	return &service{
		config:   config,
		cache:    cache,
		store:    store,
		uuidFunc: uuidFunc,
		timeFunc: timeFunc,
	}
}

type service struct {
	config   *app.Config
	cache    store.Cache
	store    Store
	uuidFunc func() uuid.UUID
	timeFunc func() time.Time
}
