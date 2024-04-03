package main

import (
	"eazimation-backend/internal/server"
	"fmt"
	"log/slog"
)

func main() {

	s := server.NewServer()
	slog.Info("Server listening", "port", s.Addr[1:])
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
