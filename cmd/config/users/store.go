package users

import (
	"database/sql"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services/users"
	appStore "github.com/sbuttigieg/test-xm/xm_app/store"
	usersStore "github.com/sbuttigieg/test-xm/xm_app/store/psql/users"
	"github.com/sbuttigieg/test-xm/xm_app/store/psql/users/middleware"
)

func NewStore(cfg *app.Config, db *sql.DB, cache appStore.Cache) users.Store {
	s := usersStore.New(db)
	s = middleware.NewCacheMiddleware(cfg, s, cache)

	return s
}
