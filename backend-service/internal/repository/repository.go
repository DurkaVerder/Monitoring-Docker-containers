package repository

import "github.com/DurkaVerder/models"

type Repository interface {
	GetPing(IPAdress string) (*models.PingResult, error)
	GetAllPing() ([]models.PingResult, error)
	AddPing(ping models.PingResult) error
	UpdatePing(ping models.PingResult) error
	Close()
}
