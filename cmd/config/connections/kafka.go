package connections

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewKafka() (redis.UniversalClient, error) {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{redisAddr},
		Password: "",
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
// 	config := sarama.NewConfig()
// 	config.Producer.Return.Successes = true
// 	config.Producer.RequiredAcks = sarama.WaitForAll
// 	config.Producer.Retry.Max = 5
// 	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
// 	conn, err := sarama.NewSyncProducer(brokersUrl, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conn, nil
// }
