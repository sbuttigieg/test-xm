package users

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/models"
)

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store
type Store interface {
	Create(context.Context, *models.User) (string, error)
	GetByEmail(context.Context, string) (*models.User, error)
}
