package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var Client *redis.Client

type IRedis interface {
	SetValue()
	GetValue()
}

type Redis struct{}

func (r Redis) NewRedisRepository() IRedis {
	return &Redis{}
}

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

func (r Redis) SetValue() {}

func (r Redis) GetValue() {}
