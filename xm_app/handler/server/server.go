package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-xm/xm_app/handler/middleware"
	"github.com/sirupsen/logrus"
)

// StartServer starts health server
func StartServer(apiAddr string, ctx context.Context, log *logrus.Logger) *http.Server {
	// Comment for debug mode. Uncomment for production
	// gin.SetMode(gin.ReleaseMode)

	// Create a new instance of the Gin router
	appRouter := gin.New()
	appRouter.Use(gin.Recovery())
	appRouter.Use(middleware.Logger(ctx, log))

	err := appRouter.SetTrustedProxies(nil)
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

	// hmux := http.NewServeMux()
	// hmux.HandleFunc("/healthz", Healthz)
	// hmux.HandleFunc("/healthz/status", HealthzStatus)
	// hmux.HandleFunc("/readiness", Readiness)
	// hmux.HandleFunc("/readiness/status", ReadinessStatus)

	// healthServer := manners.NewServer()
	// healthServer.Addr = healthAddr
	// healthServer.Handler = hmux

	return healthServer
}
