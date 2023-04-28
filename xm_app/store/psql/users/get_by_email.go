package users

import (
	"context"
	"database/sql"
	goErrors "errors"

	"github.com/lib/pq"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *store) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := s.db.QueryRowContext(ctx,
		`SELECT * FROM users WHERE email=$1`, email).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("not found user by %s", email)
		}

		var postgresError *pq.Error
		if goErrors.As(err, &postgresError) {
			err = errors.Wrapf(err, " database error: %s", postgresError.Code.Class().Name())
		}

		return nil, errors.Wrapf(err, "get user by %s", email)
	}

	return &user, nil
}
