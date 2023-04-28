package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
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
	_, err = connections.NewKafka()
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}

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

	// Configure the Kafka producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Wait for only the leader to acknowledge
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	// Connect to the Kafka broker(s)
	// brokers := []string{"localhost:9092", "localhost:29092"}
	// producer, err := sarama.NewAsyncProducer(brokers, config)
	// if err != nil {
	// 	log.WithContext(ctx).Panic(err.Error())
	// }
	// defer producer.AsyncClose()

	// // Create a Kafka message
	// msg := &sarama.ProducerMessage{
	// 	Topic: "quickstart",
	// 	Value: sarama.StringEncoder("Hello, Kafka!"),
	// }

	// // Send the message to Kafka
	// producer.Input() <- msg

	// // Wait for the message to be sent
	// select {
	// case <-producer.Successes():
	// 	fmt.Println("Message sent to Kafka")
	// case err := <-producer.Errors():
	// 	fmt.Println("Failed to send message to Kafka:", err)
	// }

	// Start the server
	err = appRouter.Run(fmt.Sprintf(":%s", apiAddr))
	if err != nil {
		log.WithContext(ctx).Panic(err.Error())
	}
}
