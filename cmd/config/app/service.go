package app

import (
	"time"

	"github.com/google/uuid"
	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services"
	"github.com/sbuttigieg/test-xm/xm_app/services/middleware"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

func NewService(cfg *app.Config, cache store.Cache, store services.Store, uuidFunc func() uuid.UUID, timeFunc func() time.Time) services.Service {
	service := services.New(cfg, cache, store, uuidFunc, timeFunc)
	service = middleware.NewLoggingMiddleware(cfg, service)

	return service
}
