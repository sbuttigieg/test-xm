package services

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
const (
	topic = "quickstart"
	// broker1Address = ""
)

type Broker struct {
	host string
	port int
}

func (b Broker) String() string {
	return fmt.Sprintf("%v:%v", b.host, b.port)
}

func Produce(ctx context.Context, msg string, key int) (bool, error) {
	broker := Broker{
		host: "localhost",
		port: 9092}

	fmt.Printf("Using Broker: %v\n--------------------------\n\n", broker)

	c := kafka.ConfigMap{
		"bootstrap.servers": broker.String()}

	// initialize a counter
	// i := 0
	// if e := doAdmin(broker); e != nil {
	// 	fmt.Printf("❌There was a problem calling the Admin Client:\n%v\n", e)
	// } else {
	// 	// Produce message
	// 	fmt.Printf("✅ AdminClient worked\n--------------------------\n\n")
	// 	if m, e := doProduce(broker, topic); e != nil {
	// 		fmt.Printf("❌There was a problem calling the producer:\n%v\n", e)
	// 	} else {
	// 		fmt.Printf("✅ Producer worked\n--------------------------\n\n")
	// 		// Consume message
	// 		if e := doConsume(broker, topic, m); e != nil {
	// 			fmt.Printf("❌There was a problem calling the consumer:\n%v\n", e)
	// 		} else {
	// 			fmt.Printf("✅ Consumer worked\n--------------------------\n\n")

	// 			// Consume message

	// 		}
	// 	}

	// }
	// intialize the writer with the broker addresses, and the topic
	w2 := kafka.Writer{
		Addr:     kafka.TCP("broker:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	// defer w2.Close()

	// msgw2 := kafka.Message{
	// 	Key:   []byte(fmt.Sprintf("address-%v", key)),
	// 	Value: []byte(msg),
	// }

	// err := w2.WriteMessages(ctx, msgw2)

	// if err != nil {
	// 	return false, errors.Wrapf(err, "kafka error")
	// }
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{"broker:9092", "broker:29092"},
	// 	Topic:   topic,
	// })

	// for {
	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	// err := w.WriteMessages(ctx, kafka.Message{
	// 	Key: []byte(strconv.Itoa(key)),
	// 	// create an arbitrary message payload for the value
	// 	Value: []byte(msg),
	// })
	// if err != nil {
	// 	return false, errors.Wrapf(err, "create user")
	// }

	// log a confirmation once the message is written
	// fmt.Println("writes:", i)
	// i++
	// sleep for a second
	// time.Sleep(time.Second * 5)
	// }

	return true, nil
}
