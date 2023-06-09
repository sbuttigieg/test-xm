package middleware

import (
	"context"
	"time"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
)

type loggingMiddleware struct {
	config *app.Config
	next   companies.Service
}

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(cfg *app.Config, next companies.Service) companies.Service {
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

func (m *loggingMiddleware) Delete(ctx context.Context, req string) error {
	start := time.Now()
	err := m.next.Delete(ctx, req)
	end := time.Now()
	m.config.Log.Infof(
		"service call, duration: %v, service-name: company, method: Get, layer: service, req: %s, resp: n/a, error: %v",
		end.Sub(start).String(), req, err,
	)

	return err
}

func (m *loggingMiddleware) Update(ctx context.Context, id string, data *models.Company, fields []string) (*models.Company, error) {
	start := time.Now()
	resp, err := m.next.Update(ctx, id, data, fields)
	end := time.Now()
	m.config.Log.Infof(
		"service call, duration: %v, service-name: company, method: Update, layer: service, req: %+v, fields: %+v, resp: %+v, error: %v",
		end.Sub(start).String(), data, fields, resp, err,
	)

	return resp, err
}
