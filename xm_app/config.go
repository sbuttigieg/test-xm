package app

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	CacheExpiry  time.Duration
	Env          string
	Log          *logrus.Logger
	ServiceName  string
	StoreTimeout time.Duration
	Version      string
}
