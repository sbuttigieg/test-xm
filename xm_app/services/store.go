package services

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/models"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	Create(context.Context, *models.Company) (string, error)
	Get(context.Context, string) (*models.Company, error)
	Update(context.Context, string, *models.Company, []string) (*models.Company, error)
}
