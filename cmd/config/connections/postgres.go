package connections

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	app "github.com/sbuttigieg/test-xm/xm_app"
)

func NewPostgres(c *app.Config) (*gorm.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_NAME")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	dbUser := os.Getenv("POSTGRES_USER")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPwd, dbName, dbPort)

	var db *gorm.DB

	var dbOK bool

	var err error

	time.Sleep(c.StoreTimeout) // time for postgres to load

	for i := 0; i <= 3; i++ {
		if !dbOK {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				// log.Info(fmt.Sprintf("DB load trial no. %v: ", i+1), err)
				time.Sleep(c.StoreTimeout) // time for retries

				continue
			}
		}

		dbOK = true

		break
	}

	return db, nil
}
