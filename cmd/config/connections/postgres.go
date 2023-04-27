package connections

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	app "github.com/sbuttigieg/test-xm/xm_app"
)

func NewPostgres(c *app.Config) (*sql.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("database host is empty")
	}

	dbName := os.Getenv("POSTGRES_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("database name is empty")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("database port is empty")
	}

	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	if dbPwd == "" {
		return nil, fmt.Errorf("database password is empty")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("database user is empty")
	}

	sqltrace.Register("postgres", pq.Driver{})

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPwd, dbName, dbPort)

	db, err := retryConn(c, dsn)
	if err != nil {
		return nil, fmt.Errorf("store connection error: %w", err)
	}

	err = runMigrations(c, db)
	if err != nil {
		return nil, fmt.Errorf("migrate error: %w", err)
	}

	return db, nil
}

func retryConn(c *app.Config, dsn string) (*sql.DB, error) {
	for i := 0; i <= 3; i++ {
		db, err := sqltrace.Open("postgres", dsn, sqltrace.WithServiceName(c.ServiceName))
		if err != nil {
			c.Log.Info(fmt.Sprintf("DB load try no. %v: ", i+1), err)
			time.Sleep(c.StoreTimeout)

			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		time.Sleep(c.StoreTimeout)
	}

	return nil, fmt.Errorf("database connection retry exceded")
}

func runMigrations(c *app.Config, db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	migrate.SetTable("migrations")

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("base migrations: %w", err)
	}

	c.Log.Info(fmt.Sprintf("Applied base %d migrations!", n))

	return nil
}
