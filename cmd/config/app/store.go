package app

import (
	"database/sql"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services"
	appStore "github.com/sbuttigieg/test-xm/xm_app/store"
	store "github.com/sbuttigieg/test-xm/xm_app/store/psql"
	"github.com/sbuttigieg/test-xm/xm_app/store/psql/middleware"
)

func NewStore(cfg *app.Config, db *sql.DB, cache appStore.Cache) services.Store {
	s := store.New(db)
	s = middleware.NewCacheMiddleware(cfg, s, cache)

	return s
}
