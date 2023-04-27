package config

import (
	"os"
	"strconv"
	"time"

	app "github.com/sbuttigieg/test-xm/xm_app"
)

// NewConfig create new config
func NewConfig() (*app.Config, error) {
	env := os.Getenv("ENV")
	version := os.Getenv("VERSION")

	cacheExpiry, err := strconv.Atoi(os.Getenv("REDIS_EXPIRY_SEC"))
	if err != nil {
		return nil, err
	}

	storeTimeout, err := strconv.Atoi(os.Getenv("MYSQL_TIMEOUT_SEC"))
	if err != nil {
		return nil, err
	}

	c := &app.Config{
		CacheExpiry:  time.Duration(cacheExpiry) * time.Second,
		Env:          env,
		StoreTimeout: time.Duration(storeTimeout) * time.Second,
		Version:      version,
	}

	return c, nil
}