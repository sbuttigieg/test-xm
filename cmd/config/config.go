package config

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
)

// NewConfig create new config
func NewConfig(log *logrus.Logger) (*app.Config, error) {
	cacheExpiry, err := strconv.Atoi(os.Getenv("REDIS_EXPIRY_SEC"))
	if err != nil {
		return nil, err
	}

	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.Wrap("env is empty")
	}

	jwtExpiry, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_SEC"))
	if err != nil {
		return nil, err
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		return nil, errors.Wrap("service name is empty")
	}

	storeTimeout, err := strconv.Atoi(os.Getenv("POSTGRES_TIMEOUT_SEC"))
	if err != nil {
		return nil, err
	}

	version := os.Getenv("VERSION")
	if version == "" {
		return nil, errors.Wrap("version is empty")
	}

	c := &app.Config{
		CacheExpiry:  time.Duration(cacheExpiry) * time.Second,
		Env:          env,
		JWTExpiry:    time.Duration(jwtExpiry) * time.Second,
		Log:          log,
		ServiceName:  serviceName,
		StoreTimeout: time.Duration(storeTimeout) * time.Second,
		Version:      version,
	}

	return c, nil
}
