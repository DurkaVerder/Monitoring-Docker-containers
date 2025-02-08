package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DurkaVerder/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetAllPing() ([]models.PingResult, error) {
	args := m.Called()
	return args.Get(0).([]models.PingResult), args.Error(1)
}

func (m *MockService) UpdateTablePings(ping models.PingResult) error {
	args := m.Called(ping)
	return args.Error(0)
}

func (m *MockService) Close() {
	m.Called()
}

func TestHandlerGetAllPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	h := NewHandlersManager(mockService)

	now := time.Now()
	expectedPings := []models.PingResult{
		{
			IPAddress: "192.168.1.1",
			PingTime:  10,
			DateSuccessfulPing: pq.NullTime{
				Time:  now,
				Valid: true,
			},
		},
	}

	tests := []struct {
		name         string
		mockReturn   []models.PingResult
		mockErr      error
		expectedCode int
	}{
		{
			name:         "Success",
			mockReturn:   expectedPings,
			mockErr:      nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Service error",
			mockReturn:   nil,
			mockErr:      errors.New("service error"),
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			mockService.ExpectedCalls = nil
			mockService.On("GetAllPing").Return(tt.mockReturn, tt.mockErr)

			h.HandlerGetAllPing(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusOK {
				var response []models.PingResult
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.mockReturn), len(response))
				
				for i, ping := range response {
					assert.Equal(t, tt.mockReturn[i].IPAddress, ping.IPAddress)
					assert.Equal(t, tt.mockReturn[i].PingTime, ping.PingTime)
					assert.Equal(t, tt.mockReturn[i].DateSuccessfulPing.Valid, ping.DateSuccessfulPing.Valid)
				}
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestHandlerAddPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	h := NewHandlersManager(mockService)

	now := time.Now()
	testPing := models.PingResult{
		IPAddress: "192.168.1.1",
		PingTime:  10,
		DateSuccessfulPing: pq.NullTime{
			Time:  now,
			Valid: true,
		},
	}

	tests := []struct {
		name         string
		input        interface{}
		mockErr      error
		expectedCode int
	}{
		{
			name:         "Success",
			input:        testPing,
			mockErr:      nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid input",
			input:        "invalid",
			mockErr:      nil,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Service error",
			input:        testPing,
			mockErr:      errors.New("service error"),
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonData, err := json.Marshal(tt.input)
			assert.NoError(t, err)

			c.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			if tt.input == testPing {
				mockService.ExpectedCalls = nil
				mockService.On("UpdateTablePings", mock.MatchedBy(func(p models.PingResult) bool {
					return p.IPAddress == testPing.IPAddress &&
						p.PingTime == testPing.PingTime &&
						p.DateSuccessfulPing.Valid == testPing.DateSuccessfulPing.Valid
				})).Return(tt.mockErr)
			}

			h.HandlerAddPing(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}
