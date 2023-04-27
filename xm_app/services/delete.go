package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
)

func (s *service) Delete(ctx context.Context, id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.BadRequest("id is not a valid uuid")
	}

	err = s.store.Delete(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "delete company")
	}

	return nil
}
