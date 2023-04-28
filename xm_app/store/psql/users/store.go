package users

import (
	"database/sql"

	"github.com/sbuttigieg/test-xm/xm_app/services/users"
)

func New(db *sql.DB) users.Store {
	s := &store{
		db:     db,
		models: "company",
	}

	return s
}

type store struct {
	db     *sql.DB
	models string
}
