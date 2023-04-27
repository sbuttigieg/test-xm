package connections

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	app "github.com/sbuttigieg/test-xm/xm_app"
)

func NewMySQL(c *app.Config) (*gorm.DB, error) {
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbPort := os.Getenv("MYSQL_PORT")
	dbPwd := os.Getenv("MYSQL_PASSWORD")
	dbUser := os.Getenv("MYSQL_USER")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)

	var db *gorm.DB

	var dbOK bool

	var err error

	time.Sleep(c.StoreTimeout) // time for mysql to load

	for i := 0; i <= 3; i++ {
		if !dbOK {
			db, err = gorm.Open(mysql.Open(dsn))
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
