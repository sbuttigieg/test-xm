package companies

import (
	"time"

	"github.com/google/uuid"
	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies/middleware"
	"github.com/sbuttigieg/test-xm/xm_app/store"
)

func NewService(cfg *app.Config, cache store.Cache, store companies.Store, uuidFunc func() uuid.UUID, timeFunc func() time.Time) companies.Service {
	service := companies.New(cfg, cache, store, uuidFunc, timeFunc)
	service = middleware.NewLoggingMiddleware(cfg, service)

	return service
}
