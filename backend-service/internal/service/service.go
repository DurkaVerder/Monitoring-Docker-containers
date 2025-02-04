package service

import "backend-service/internal/repository"

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

