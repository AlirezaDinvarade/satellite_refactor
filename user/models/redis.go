package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	REDIS_HOST     = os.Getenv("REDIS_HOST")
	REDIS_PORT     = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	RedisDB        *redis.Client
)

type RedisAdaptor struct {
	Client *redis.Client
}

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	SetEx(ctx context.Context, key string, value []byte, TTL time.Duration) error
	Del(ctx context.Context, key string) error
}

func (r *RedisAdaptor) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisAdaptor) SetEx(ctx context.Context, key string, value []byte, TTL time.Duration) error {
	return r.Client.SetEx(ctx, key, value, TTL).Err()
}

func (r *RedisAdaptor) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}


func ConnectRedis() {
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
