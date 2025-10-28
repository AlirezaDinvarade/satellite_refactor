package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func ConnectRedis() {
	var (
		REDIS_HOST     = os.Getenv("REDIS_HOST")
		REDIS_PORT     = os.Getenv("REDIS_PORT")
		REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	)

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Password: REDIS_PASSWORD,
		DB:       0,
	})

	_, err := RedisDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("redis Connected")
}
