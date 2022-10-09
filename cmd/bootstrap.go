package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffr-k/waffletime/internal/auth"
	"github.com/jeffr-k/waffletime/internal/user/presentor"
	"github.com/jeffr-k/waffletime/pkg/database"
	"github.com/segmentio/kafka-go"
)

type Bootstrap struct{}

func (b Bootstrap) initializeRouter(router *gin.Engine) {
	user := presentor.Router{}
	auth := auth.Router{}

	user.Routes(router)
	auth.Routes(router)
}

func (b Bootstrap) initializedDatabase() {
	database.InitializedDatabase()
}

func (b Bootstrap) initializeKafka() {
	topic := "waffletime"
	partition := 0

	_, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		panic("Kafka connection is failed.")
	}
	fmt.Println("kafka connection is success..")

}

func (b Bootstrap) initializeMongoDB() {}
