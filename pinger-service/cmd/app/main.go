package main

import (
	"context"
	"log"
	"os"
	"pinger-service/config"
	"pinger-service/internal/pinger"

	"github.com/DurkaVerder/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	config := config.LoadConfig(os.Getenv("CONFIG_PATH"))

	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}
	pingResultChan := make(chan models.PingResult, config.Channel.Size)
	dockerContainers := make(chan types.Container, config.Channel.Size)

	pingerService := pinger.NewPingerService(client, pingResultChan, dockerContainers, config)

	ctx := context.Background()

	go pingerService.Run(ctx)

	signalChan := make(chan os.Signal, 1)

	sig := <-signalChan

	log.Printf("Received signal: %v. Shutting down", sig)

	pingerService.Stop()
}
