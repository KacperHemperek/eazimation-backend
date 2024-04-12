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
	redisHost     = os.Getenv("REDIS_HOST")
)

func NewRedisClient() *redis.Client {
	slog.Info("redis values", "host", redisHost, "address", redisAddr)
	redisConnStr := fmt.Sprintf("%s:%s", redisHost, redisAddr)
	slog.Info("Redis address", "address", redisConnStr)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConnStr,
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
