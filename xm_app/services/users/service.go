package users

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
	Create(context.Context, *models.User) (string, error)
	GetByEmail(context.Context, string) (*models.User, error)
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
