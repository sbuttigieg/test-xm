package store

import (
	"database/sql"

	"github.com/sbuttigieg/test-xm/xm_app/services"
)

func New(db *sql.DB) services.Store {
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
