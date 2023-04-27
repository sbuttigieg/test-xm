package services

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/models"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	Get(context.Context, string) (*models.Company, error)
}
