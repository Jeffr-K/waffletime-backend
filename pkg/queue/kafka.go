package queue

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type IKafkaQueue interface {
	Produce()
	Consume()
}

type kafkaQueue struct{}

func NewKafka() IKafkaQueue {
	return &kafkaQueue{}
}

func (k kafkaQueue) initializeKafka() {
	topic := "waffletime"
	partition := 0

	_, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatalln("Kafka connection Error: ", err)
		panic(err)
	}
	fmt.Println("kafka connection is success..")
}

func (k kafkaQueue) Produce() {
	writer := &kafka.Writer{
		Addr:        kafka.TCP("localhost:9092"),
		Topic:       "user.registered",
		Balancer:    &kafka.LeastBytes{},
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("userId"),
		Value: []byte("user.registered"),
	})
	if err != nil {
		log.Fatalln("failed to write messages", err)
	}
	if err := writer.Close(); err != nil {
		log.Fatalln("failed to close writer: ", err)
	}
}

func (k kafkaQueue) Consume() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "user.registered",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	reader.SetOffset(42)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader: ", err)
	}

}

func logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}
