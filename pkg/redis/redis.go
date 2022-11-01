package redis

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v9"
	"log"
)

var Client *redis.Client

type IRedis interface {
	RedisClientConnection() (*redis.Client, error)
}

type Redis struct{}

func (r Redis) NewRedisRepository() IRedis {
	return &Redis{}
}

func InitializeRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("\n[Redis Connection Error]\n", err)
		panic(err)
	}

	fmt.Println("\n[Redis Connection Success]\n", pong)
}

func (r Redis) RedisClientConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("\n[Redis Connection Error]\n", err)
		panic(err)
	}

	return client, nil
}