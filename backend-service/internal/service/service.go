package service

import (
	"backend-service/internal/repository"
	"log"

	"github.com/DurkaVerder/models"
)

type Service interface {
}

type ServiceManager struct {
	repo repository.Repository
}

func NewServiceManager(repo repository.Repository) *ServiceManager {
	return &ServiceManager{
		repo: repo,
	}
}

func (s *ServiceManager) Close() {
	s.repo.Close()
}

func (s *ServiceManager) GetAllPing() ([]models.PingResult, error) {
	pings, err := s.repo.GetAllPing()
	if err != nil {
		log.Printf("Error getting all pings: %v", err)
		return nil, err
	}

	return pings, nil
}

func (s *ServiceManager) AddPing(ping models.PingResult) error {
	if err := s.repo.AddPing(ping); err != nil {
		log.Printf("Error adding ping: %v", err)
	}

	return nil
}
