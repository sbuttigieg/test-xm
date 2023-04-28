package companies

import (
	"database/sql"

	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
)

func New(db *sql.DB) companies.Store {
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
