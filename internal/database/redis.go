package database

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
)

var (
	redisAddr     = os.Getenv("REDIS_PORT")
	redisPassword = os.Getenv("REDIS_PASSWORD")
)

func NewRedisClient() *redis.Client {
	slog.Info("Redis address", "address", fmt.Sprintf("localhost:%s", redisAddr))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("127.0.0.1:%s", redisAddr),
		Password: redisPassword,
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Info(ctx).Result()
	if err != nil {
		panic(err)
	}
	return client
}
