package repository

import "github.com/DurkaVerder/models"

type Repository interface {
	GetAllPing() ([]models.PingResult, error)
	AddPing(ping models.PingResult) error
}


