package app

import (
	"time"
)

type Config struct {
	CacheExpiry  time.Duration
	Env          string
	ServiceName  string
	StoreTimeout time.Duration
	Version      string
}
