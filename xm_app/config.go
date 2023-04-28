package app

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	CacheExpiry  time.Duration
	Env          string
	JWTExpiry    time.Duration
	Log          *logrus.Logger
	ServiceName  string
	StoreTimeout time.Duration
	Version      string
}
