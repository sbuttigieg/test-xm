package users

import (
	"time"

	"github.com/google/uuid"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services/users"
	"github.com/sbuttigieg/test-xm/xm_app/services/users/middleware"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

func NewService(cfg *app.Config, cache store.Cache, store users.Store, uuidFunc func() uuid.UUID, timeFunc func() time.Time) users.Service {
	service := users.New(cfg, cache, store, uuidFunc, timeFunc)
	service = middleware.NewLoggingMiddleware(cfg, service)

	return service
}
