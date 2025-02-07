package service

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/DurkaVerder/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRepository struct {
// 	mock.Mock
// }

// func (m *MockRepository) GetAllPing() ([]models.PingResult, error) {
// 	args := m.Called()
// 	return args.Get(0).([]models.PingResult), args.Error(1)
// }

// func (m *MockRepository) GetPing(IPAdress string) (*models.PingResult, error) {
// 	args := m.Called(IPAdress)
// 	return args.Get(0).(*models.PingResult), args.Error(1)
// }

// func (m *MockRepository) AddPing(ping models.PingResult) error {
// 	args := m.Called(ping)
// 	return args.Error(0)
// }

// func (m *MockRepository) UpdatePing(ping models.PingResult) error {
// 	args := m.Called(ping)
// 	return args.Error(0)
// }

// func (m *MockRepository) Close() {
// 	m.Called()
// }

// func TestGetAllPing(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewServiceManager(mockRepo)

// 	date := time.Now()
// 	expectedErr := errors.New("Error connection to database")

// 	tests := []struct {
// 		name    string
// 		mockRet []models.PingResult
// 		mockErr error
// 		want    []models.PingResult
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Successful retrieval of ping results",
// 			mockRet: []models.PingResult{{IPAddress: "192.12.31.3", PingTime: 12, DateSuccessfulPing: &date}},
// 			mockErr: nil,
// 			want:    []models.PingResult{{IPAddress: "192.12.31.3", PingTime: 12, DateSuccessfulPing: &date}},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "No results returned",
// 			mockRet: []models.PingResult(nil),
// 			mockErr: nil,
// 			want:    []models.PingResult(nil),
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Database error",
// 			mockRet: nil,
// 			mockErr: expectedErr,
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockRepo.ExpectedCalls = nil
// 			mockRepo.On("GetAllPing").Return(tt.mockRet, tt.mockErr)

// 			got, err := s.GetAllPing()

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.mockErr, err)
// 				assert.Nil(t, got)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.want, got)
// 			}

// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }

// func TestUpdateTablePings(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	s := NewServiceManager(mockRepo)

// 	date := time.Now()
// 	testPing := models.PingResult{
// 		IPAddress:          "192.168.1.1",
// 		PingTime:           10,
// 		DateSuccessfulPing: date,
// 	}

// 	tests := []struct {
// 		name          string
// 		mockGetRet    *models.PingResult
// 		mockGetErr    error
// 		mockAddErr    error
// 		mockUpdateErr error
// 		wantErr       bool
// 	}{
// 		{
// 			name:          "Ping does not exist, add successfully",
// 			mockGetRet:    nil,
// 			mockGetErr:    nil,
// 			mockAddErr:    nil,
// 			mockUpdateErr: nil,
// 			wantErr:       false,
// 		},
// 		{
// 			name:          "Ping does not exist, add fails",
// 			mockGetRet:    nil,
// 			mockGetErr:    nil,
// 			mockAddErr:    errors.New("failed to add ping"),
// 			mockUpdateErr: nil,
// 			wantErr:       true,
// 		},
// 		{
// 			name:          "Ping exists, update successfully",
// 			mockGetRet:    &testPing,
// 			mockGetErr:    nil,
// 			mockAddErr:    nil,
// 			mockUpdateErr: nil,
// 			wantErr:       false,
// 		},
// 		{
// 			name:          "Ping exists, update fails",
// 			mockGetRet:    &testPing,
// 			mockGetErr:    nil,
// 			mockAddErr:    nil,
// 			mockUpdateErr: errors.New("failed to update ping"),
// 			wantErr:       true,
// 		},
// 		{
// 			name:          "GetPing fails",
// 			mockGetRet:    nil,
// 			mockGetErr:    errors.New("failed to get ping"),
// 			mockAddErr:    nil,
// 			mockUpdateErr: nil,
// 			wantErr:       true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockRepo.ExpectedCalls = nil
// 			mockRepo.On("GetPing", testPing.IPAddress).Return(tt.mockGetRet, tt.mockGetErr)

// 			if tt.mockGetRet == nil && tt.mockGetErr == nil {
// 				mockRepo.On("AddPing", testPing).Return(tt.mockAddErr)
// 			}

// 			if tt.mockGetRet != nil && tt.mockGetErr == nil {
// 				mockRepo.On("UpdatePing", testPing).Return(tt.mockUpdateErr)
// 			}

// 			err := s.UpdateTablePings(testPing)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}

// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }
