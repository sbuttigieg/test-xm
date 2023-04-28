package users

import (
	"context"
	goErrors "errors"

	"github.com/lib/pq"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *store) Create(ctx context.Context, req *models.User) (string, error) {
	var id string

	//nolint
	err := s.db.QueryRowContext(ctx, `INSERT INTO users
	(name, username, email, password)
	VALUES ( $1, $2, $3, $4) RETURNING id`,
		req.Name,
		req.Username,
		req.Email,
		req.Password,
	).Scan(&id)
	if err != nil {
		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return "", errors.Wrapf(err, "create user")
	}

	return id, nil
}
