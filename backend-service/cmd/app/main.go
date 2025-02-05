package main

import (
	"backend-service/internal/handlers"
	"backend-service/internal/repository/postgres"
	"backend-service/internal/server"
	"backend-service/internal/service"
)

func main() {
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
	server := server.NewServer(handlers)

	// Run server
	server.Run(":8080")

}
