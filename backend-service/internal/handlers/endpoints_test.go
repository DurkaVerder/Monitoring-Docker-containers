package handlers

// import (
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/DurkaVerder/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockService struct {
// 	mock.Mock
// }

// func (m *MockService) GetAllPing() ([]models.PingResult, error) {
// 	args := m.Called()
// 	return args.Get(0).([]models.PingResult), args.Error(1)
// }

// func (m *MockService) UpdateTablePings(newPing models.PingResult) error {
// 	args := m.Called(newPing)
// 	return args.Error(0)
// }

// func (m *MockService) Close() {
// 	m.Called()
// }

// func TestHandlerGetAllPing(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockService := MockService{}
// 	h := HandlersManager{&mockService}

// 	date := time.Now()
// 	pingResults := []models.PingResult{
// 		{IPAddress: "192.168.1.1", PingTime: 10, DateSuccessfulPing: &date},
// 		{IPAddress: "192.168.1.2", PingTime: 15, DateSuccessfulPing: &date},
// 	}

// 	tests := []struct {
// 		name           string
// 		mockGetAllRet  []models.PingResult
// 		mockGetAllErr  error
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "Success - multiple results",
// 			mockGetAllRet:  pingResults,
// 			mockGetAllErr:  nil,
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `[{"ip_address":"192.168.1.1","ping_time":10,"date_successful_ping":"` + date.Format("2006-01-02T15:04:05.0000000Z07:00") + `","id":0},{"ip_address":"192.168.1.2","ping_time":15,"date_successful_ping":"` + date.Format("2006-01-02T15:04:05.0000000Z07:00") + `","id":0}]`,
// 		},
// 		{
// 			name:           "Error - service fails",
// 			mockGetAllRet:  nil,
// 			mockGetAllErr:  errors.New("database error"),
// 			expectedStatus: http.StatusInternalServerError,
// 			expectedBody:   `{"error":"database error"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockService.ExpectedCalls = nil

// 			mockService.On("GetAllPing").Return(tt.mockGetAllRet, tt.mockGetAllErr)

// 			w := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(w)

// 			h.HandlerGetAllPing(c)

// 			assert.Equal(t, tt.expectedStatus, w.Code)

// 			assert.JSONEq(t, tt.expectedBody, w.Body.String())

// 			mockService.AssertExpectations(t)
// 		})
// 	}
// }
