package main

import (
	"context"
	"log"
	"os"

	"github.com/sbuttigieg/test-xm/cmd/config"
	"github.com/sbuttigieg/test-xm/cmd/config/connections"
	"github.com/sbuttigieg/test-xm/cmd/config/store"
)

func main() {
	ctx := context.Background()

	// logger setup
	logFile := "logs.txt"

	//nolint
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	log, err := config.NewLogger(f)
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	// config
	c, err := config.NewConfig(log)
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	// connections
	redisConnection, err := connections.NewRedis()
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	dbConnection, err := connections.NewPostgres(c)
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	// redis setup
	cache := store.NewCache(c, redisConnection)

	log.WithContext(ctx).Info(dbConnection, cache)
}
