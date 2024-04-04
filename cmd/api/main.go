package main

import (
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/server"
	"fmt"
	"log/slog"
)

func main() {
	auth.NewAuth()
	s := server.NewServer()
	slog.Info("Server listening", "port", s.Addr[1:])
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
