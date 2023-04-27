package middleware

import (
	"context"
	"time"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services"
)

type loggingMiddleware struct {
	config *app.Config
	next   services.Service
}

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(cfg *app.Config, next services.Service) services.Service {
	m := loggingMiddleware{
		config: cfg,
		next:   next,
	}

	return &m
}

func (m *loggingMiddleware) Create(ctx context.Context, req *models.Company) (string, error) {
	start := time.Now()
	resp, err := m.next.Create(ctx, req)
	end := time.Now()
	m.config.Log.Infof(
		"service call, duration: %v, service-name: company, method: Create, layer: service, req: %+v, resp: %s, error: %v",
		end.Sub(start).String(), req, resp, err,
	)

	return resp, err
}

func (m *loggingMiddleware) Get(ctx context.Context, req string) (*models.Company, error) {
	start := time.Now()
	resp, err := m.next.Get(ctx, req)
	end := time.Now()
	m.config.Log.Infof(
		"service call, duration: %v, service-name: company, method: Get, layer: service, req: %s, resp: %+v, error: %v",
		end.Sub(start).String(), req, resp, err,
	)

	return resp, err
}
