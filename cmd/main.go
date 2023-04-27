package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sbuttigieg/test-xm/cmd/config"
	"github.com/sbuttigieg/test-xm/cmd/config/app"
	"github.com/sbuttigieg/test-xm/cmd/config/connections"
	"github.com/sbuttigieg/test-xm/cmd/config/store"
	"github.com/sbuttigieg/test-xm/xm_app/handler/middleware"
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

	// api setup
	endpointURL := os.Getenv("ENDPOINT_URL")
	apiAddr := os.Getenv("PORT")
	appStore := app.NewStore(c, dbConnection, cache)
	appService := app.NewService(c, cache, appStore, uuid.New, time.Now)
	appHandlers := app.NewHandlers(appService)

	// Comment for debug mode. Uncomment for production
	// gin.SetMode(gin.ReleaseMode)

	// Create a new instance of the Gin router
	appRouter := gin.New()
	appRouter.Use(gin.Recovery())
	appRouter.Use(middleware.Logger(ctx, log))

	err = appRouter.SetTrustedProxies(nil)
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	// Endpoints
	appRouter.POST(endpointURL, appHandlers.Create)
	appRouter.DELETE(fmt.Sprintf("%s/:id", endpointURL), appHandlers.Delete)
	appRouter.GET(fmt.Sprintf("%s/:id", endpointURL), appHandlers.Get)
	appRouter.PATCH(fmt.Sprintf("%s/:id", endpointURL), appHandlers.Update)
	// middleware.BasicAuth(appService)

	// Start the server
	err = appRouter.Run(fmt.Sprintf(":%s", apiAddr))
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}
}
