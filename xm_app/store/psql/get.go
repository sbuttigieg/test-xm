package store

import (
	"context"
	"database/sql"
	goErrors "errors"

	"github.com/lib/pq"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *store) Get(ctx context.Context, id string) (*models.Company, error) {
	var company models.Company

	err := s.db.QueryRowContext(ctx,
		`SELECT * FROM companies WHERE id=$1`, id).Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.Employees,
		&company.Registered,
		&company.Type,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("not found company by %s", id)
		}

		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return nil, errors.Wrapf(err, "get company by %s", id)
	}

	return &company, nil
}
