package service

import (
	"backend-service/internal/repository"
	"log"

	"github.com/DurkaVerder/models"
)

type Service interface {
	GetAllPing() ([]models.PingResult, error)
	UpdateTablePings(newPing models.PingResult) error
	Close()
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

func (s *ServiceManager) UpdateTablePings(newPing models.PingResult) error {
	ping, err := s.repo.GetPing(newPing.IPAddress)
	if err != nil {
		log.Printf("Error getting ping: %v", err)
		return err
	}

	if ping == nil {
		if err := s.repo.AddPing(newPing); err != nil {
			log.Printf("Error adding ping: %v", err)
			return err
		}
	} else {
		if err := s.repo.UpdatePing(newPing); err != nil {
			log.Printf("Error updating ping: %v", err)
			return err
		}
	}

	return nil
}
