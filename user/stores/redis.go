package stores

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	RedisDB *redis.Client
)

type RedisAdaptor struct {
	Client *redis.Client
}

func NewRedisAdaptor() *RedisAdaptor {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading file .env")
	}
	var (
		REDIS_HOST     = os.Getenv("REDIS_HOST")
		REDIS_PORT     = os.Getenv("REDIS_PORT")
		REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Password: REDIS_PASSWORD,
		DB:       0,
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	return &RedisAdaptor{Client: client}
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
