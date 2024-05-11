package server

import (
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/services/user"
	"fmt"
	"log"
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
	db := database.New()

	err := db.Migrate()

	if err != nil {
		log.Panic(err)
	}

	redis := database.NewRedisClient()

	// initialize services
	addProviderToContext := auth.NewAddProviderToContext()
	redisSessionStore := auth.NewRedisSession(redis)

	userService := services.NewPGUserService(db)

	// initialize middlewares
	authMiddleware := auth.NewAuthMiddleware(redisSessionStore)

	// Declare Server config
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(
			addProviderToContext,
			authMiddleware,
			redisSessionStore,
			userService,
		),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
