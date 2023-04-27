package services

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *service) Create(ctx context.Context, req *models.Company) (string, error) {
	id, err := s.store.Create(ctx, req)
	if err != nil {
		return "", errors.Wrapf(err, "create company")
	}

	return id, nil
}
