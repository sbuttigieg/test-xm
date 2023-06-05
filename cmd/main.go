package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sbuttigieg/test-xm/cmd/config"
	"github.com/sbuttigieg/test-xm/cmd/config/companies"
	"github.com/sbuttigieg/test-xm/cmd/config/connections"
	"github.com/sbuttigieg/test-xm/cmd/config/store"
	"github.com/sbuttigieg/test-xm/cmd/config/users"
	"github.com/sbuttigieg/test-xm/xm_app/handler/middleware"
)

func main() {
	const errorChan int = 10

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
	healthAddr := os.Getenv("H_PORT")
	appStore := companies.NewStore(c, dbConnection, cache)
	appService := companies.NewService(c, cache, appStore, uuid.New, time.Now)
	appHandlers := companies.NewHandlers(appService)
	usersStore := users.NewStore(c, dbConnection, cache)
	usersService := users.NewService(c, cache, usersStore, uuid.New, time.Now)
	usersHandlers := users.NewHandlers(c, usersService)

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

	// Company endpoints
	company := appRouter.Group(endpointURL)
	{
		company.GET("/:id", appHandlers.Get)
		secured := appRouter.Group(endpointURL).Use(middleware.Auth(c))
		{
			secured.POST("", appHandlers.Create)
			secured.DELETE("/:id", appHandlers.Delete)
			secured.PATCH("/:id", appHandlers.Update)
		}
	}

	// User endpoints
	user := appRouter.Group(endpointURL)
	{
		user.POST("/users", usersHandlers.Create)
		user.POST("/token", usersHandlers.GetToken)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", apiAddr),
		Handler: appRouter,
	}

	errChan := make(chan error, errorChan)

	go func() {
		log.WithContext(ctx).Info(fmt.Sprintf("HTTP service listening on %s", apiAddr))
		errChan <- srv.ListenAndServe()
	}()

	healthSrv := &http.Server{
		Addr:    fmt.Sprintf(":%s", healthAddr),
		Handler: appRouter,
	}

	go func() {
		log.Info(ctx, fmt.Sprintf("Health service listening on %s", healthAddr))
		errChan <- healthSrv.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.WithContext(ctx).Fatal(err.Error())
			}
		case s := <-signalChan:
			log.Info(ctx, "Captured %v. Exiting...", s)
			// health.SetReadinessStatus(http.StatusServiceUnavailable)
			srv.Shutdown(ctx)

			os.Exit(0)
		}
	}
}
