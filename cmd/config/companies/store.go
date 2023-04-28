package companies

import (
	"database/sql"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
	appStore "github.com/sbuttigieg/test-xm/xm_app/store"
	store "github.com/sbuttigieg/test-xm/xm_app/store/psql/companies"
	"github.com/sbuttigieg/test-xm/xm_app/store/psql/companies/middleware"
)

func NewStore(cfg *app.Config, db *sql.DB, cache appStore.Cache) companies.Store {
	s := store.New(db)
	s = middleware.NewCacheMiddleware(cfg, s, cache)

	return s
}
