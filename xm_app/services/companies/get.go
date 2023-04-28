package companies

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *service) Get(ctx context.Context, id string) (*models.Company, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.BadRequest("id is not a valid uuid")
	}

	company, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "get company")
	}

	return company, nil
}
