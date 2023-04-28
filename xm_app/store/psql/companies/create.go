package companies

import (
	"context"
	goErrors "errors"

	"github.com/lib/pq"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *store) Create(ctx context.Context, req *models.Company) (string, error) {
	var id string

	//nolint
	err := s.db.QueryRowContext(ctx, `INSERT INTO companies
	(name, description, employees, registered, type)
	VALUES ( $1, $2, $3, $4, $5) RETURNING id`,
		req.Name,
		req.Description,
		req.Employees,
		req.Registered,
		req.Type,
	).Scan(&id)
	if err != nil {
		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return "", errors.Wrapf(err, "create company")
	}

	return id, nil
}
