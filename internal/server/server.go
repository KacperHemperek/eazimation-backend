package server

import (
	"eazimation-backend/internal/auth"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"eazimation-backend/internal/database"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}
	_ = database.New()
	redis := database.NewRedisClient()

	addProviderToContext := auth.NewAddProviderToContext()
	redisSessionStore := auth.NewRedisSession(redis)

	authMiddleware := auth.NewAuthMiddleware(redisSessionStore)
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(addProviderToContext, authMiddleware, redisSessionStore),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
