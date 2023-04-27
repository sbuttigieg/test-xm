package store

import (
	"context"
	goErrors "errors"

	"github.com/lib/pq"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
)

func (s *store) Delete(ctx context.Context, id string) error {
	res, err := s.db.ExecContext(ctx,
		`DELETE FROM companies WHERE id=$1`, id)
	if err != nil {
		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return errors.Wrapf(err, "delete company by %s", id)
	}

	count, err := res.RowsAffected()
	if err != nil {
		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return errors.Wrapf(err, "delete company by %s", id)
	}

	if count == 0 {
		return errors.NotFound("not found company by %s", id)
	}

	return nil
}
