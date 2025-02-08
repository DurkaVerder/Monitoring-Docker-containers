package mocks

import (
	"github.com/DurkaVerder/models"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetPing(IPAddress string) (*models.PingResult, error) {
	args := m.Called(IPAddress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.PingResult), args.Error(1)
}

func (m *MockRepository) GetAllPing() ([]models.PingResult, error) {
	args := m.Called()
	return args.Get(0).([]models.PingResult), args.Error(1)
}

func (m *MockRepository) AddPing(ping models.PingResult) error {
	args := m.Called(ping)
	return args.Error(0)
}

func (m *MockRepository) UpdatePing(ping models.PingResult) error {
	args := m.Called(ping)
	return args.Error(0)
}

func (m *MockRepository) Close() {
	m.Called()
}
