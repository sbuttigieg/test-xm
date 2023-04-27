package app

import (
	"github.com/sbuttigieg/test-xm/xm_app/handler"
	"github.com/sbuttigieg/test-xm/xm_app/services"
)

func NewHandlers(service services.Service) *handler.Handler {
	return handler.New(service)
}
