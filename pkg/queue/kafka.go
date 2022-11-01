package queue

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type IKafkaQueue interface {
	Produce()
	Consume()
}

type KafkaQueue struct{}

//func NewKafka() IKafkaQueue {
//	return &kafkaQueue{}
//}

func InitializeKafka() {
	topic := "waffletime"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatalln("Kafka connection Error: ", err)
		panic(err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatalln("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}
}

//func (k KafkaQueue) Produce() {
//	writer := &kafka.Writer{
//		Addr:        kafka.TCP("localhost:9092"),
//		Topic:       "user.registered",
//		Balancer:    &kafka.LeastBytes{},
//		Logger:      kafka.LoggerFunc(logf),
//		ErrorLogger: kafka.LoggerFunc(logf),
//	}
//
//	err := writer.WriteMessages(context.Background(), kafka.Message{
//		Key:   []byte("userId"),
//		Value: []byte("user.registered"),
//	})
//	if err != nil {
//		log.Fatalln("failed to write messages", err)
//	}
//	if err := writer.Close(); err != nil {
//		log.Fatalln("failed to close writer: ", err)
//	}
//}
//
//func (k KafkaQueue) Consume() {
//	reader := kafka.NewReader(kafka.ReaderConfig{
//		Brokers:   []string{"localhost:9092"},
//		Topic:     "user.registered",
//		Partition: 0,
//		MinBytes:  10e3,
//		MaxBytes:  10e6,
//	})
//	reader.SetOffset(42)
//
//	for {
//		message, err := reader.ReadMessage(context.Background())
//		if err != nil {
//			break
//		}
//		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
//	}
//
//	if err := reader.Close(); err != nil {
//		log.Fatal("failed to close reader: ", err)
//	}
//
//}
//
//func logf(msg string, a ...interface{}) {
//	fmt.Printf(msg, a...)
//	fmt.Println()
//}
