package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var Client *redis.Client

type Redis struct{}

func InitializeRedis() {
	database := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := database.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("\n[Redis Connection Error]\n", err)
		panic(err)
	}

	fmt.Println("\n[Redis Connection Success]\n", pong)
}
