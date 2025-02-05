package main

import (
	"backend-service/config"
	"backend-service/internal/handlers"
	"backend-service/internal/repository/postgres"
	"backend-service/internal/server"
	"backend-service/internal/service"
	"os"
)

func main() {
	cfg := config.LoadConfig(os.Getenv("CONFIG_PATH"))

	// Init DB
	db := postgres.InitDB()

	// Init repository
	repo := postgres.NewPostgres(db)

	// Init service
	service := service.NewServiceManager(repo)
	defer service.Close()

	// Init handlers
	handlers := handlers.NewHandlersManager(service)

	// Init server
	server := server.NewServer(handlers, cfg)

	// Run server
	server.Run()

}
