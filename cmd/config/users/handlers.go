package users

import (
	app "github.com/sbuttigieg/test-xm/xm_app"
	handler "github.com/sbuttigieg/test-xm/xm_app/handler/users"
	"github.com/sbuttigieg/test-xm/xm_app/services/users"
)

func NewHandlers(cfg *app.Config, service users.Service) *handler.Handler {
	return handler.New(cfg, service)
}
