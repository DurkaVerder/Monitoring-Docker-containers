package pinger

import "github.com/DurkaVerder/models"

type Pinger interface {
	PingAllContainer() ([]models.PingResult, error)
	PingContainer() (models.PingResult, error)
}

type PingerService struct {
	// slice for docker containers
}
