package users

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *service) GetByEmail(ctx context.Context, req string) (*models.User, error) {
	user, err := s.store.GetByEmail(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "get user by email")
	}

	return user, nil
}
