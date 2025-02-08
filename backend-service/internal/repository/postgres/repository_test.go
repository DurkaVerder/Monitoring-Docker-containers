package postgres

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DurkaVerder/models"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	TestGetPingQuery            = `SELECT \* FROM ping_result WHERE ip_address = \$1`
	TestAllPingQuery            = `SELECT \* FROM ping_result`
	TestAddPingQuery            = `INSERT INTO ping_result \(ip_address, ping_time, date_successful_ping\) VALUES \(\$1, \$2, \$3\)`
	TestUpdatePingWithDateQuery = `UPDATE ping_result SET ping_time = \$2, date_successful_ping = \$3 WHERE ip_address = \$1`
	TestUpdatePingQuery         = `UPDATE ping_result SET ping_time = \$2 WHERE ip_address = \$1`
)

func TestGetAllPing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewPostgres(db)

	tests := []struct {
		name          string
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedPings []models.PingResult
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "ip_address", "ping_time", "date_successful_ping"}).
					AddRow(1, "192.168.1.1", 10, time.Now()).
					AddRow(2, "192.168.1.2", 20, time.Now())

				mock.ExpectQuery(TestAllPingQuery).
					WillReturnRows(rows)
			},
			expectedPings: []models.PingResult{
				{
					ID:        1,
					IPAddress: "192.168.1.1",
					PingTime:  10,
					DateSuccessfulPing: pq.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				},
				{
					ID:        2,
					IPAddress: "192.168.1.2",
					PingTime:  20,
					DateSuccessfulPing: pq.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "Empty Result",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "ip_address", "ping_time", "date_successful_ping"})

				mock.ExpectQuery(TestAllPingQuery).
					WillReturnRows(rows)
			},
			expectedPings: []models.PingResult{},
			expectedError: nil,
		},
		{
			name: "Database Error",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(TestAllPingQuery).
					WillReturnError(errors.New("database error"))
			},
			expectedPings: nil,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			pings, err := repo.GetAllPing()

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedPings), len(pings))
				
				for i, ping := range pings {
					assert.Equal(t, tt.expectedPings[i].ID, ping.ID)
					assert.Equal(t, tt.expectedPings[i].IPAddress, ping.IPAddress)
					assert.Equal(t, tt.expectedPings[i].PingTime, ping.PingTime)
					assert.Equal(t, tt.expectedPings[i].DateSuccessfulPing.Valid, ping.DateSuccessfulPing.Valid)
				}
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetPing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewPostgres(db)

	tests := []struct {
		name          string
		ipAddress     string
		mockBehavior  func(mock sqlmock.Sqlmock, ipAddress string)
		expectedPing  *models.PingResult
		expectedError error
	}{
		{
			name:      "Success",
			ipAddress: "192.168.1.1",
			mockBehavior: func(mock sqlmock.Sqlmock, ipAddress string) {
				rows := sqlmock.NewRows([]string{"id", "ip_address", "ping_time", "date_successful_ping"}).
					AddRow(1, ipAddress, 10, time.Now())

				mock.ExpectQuery(TestGetPingQuery).
					WithArgs(ipAddress).
					WillReturnRows(rows)
			},
			expectedPing: &models.PingResult{
				ID:        1,
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			},
			expectedError: nil,
		},
		{
			name:      "Not Found",
			ipAddress: "192.168.1.1",
			mockBehavior: func(mock sqlmock.Sqlmock, ipAddress string) {
				rows := sqlmock.NewRows([]string{"id", "ip_address", "ping_time", "date_successful_ping"})

				mock.ExpectQuery(TestGetPingQuery).
					WithArgs(ipAddress).
					WillReturnRows(rows)
			},
			expectedPing:  nil,
			expectedError: sql.ErrNoRows,
		},
		{
			name:      "Database Error",
			ipAddress: "192.168.1.1",
			mockBehavior: func(mock sqlmock.Sqlmock, ipAddress string) {
				mock.ExpectQuery(TestGetPingQuery).
					WithArgs(ipAddress).
					WillReturnError(errors.New("database error"))
			},
			expectedPing:  nil,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock, tt.ipAddress)

			ping, err := repo.GetPing(tt.ipAddress)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPing.ID, ping.ID)
				assert.Equal(t, tt.expectedPing.IPAddress, ping.IPAddress)
				assert.Equal(t, tt.expectedPing.PingTime, ping.PingTime)
				assert.Equal(t, tt.expectedPing.DateSuccessfulPing.Valid, ping.DateSuccessfulPing.Valid)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAddPing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewPostgres(db)

	now := time.Now()
	tests := []struct {
		name          string
		ping          models.PingResult
		mockBehavior  func(mock sqlmock.Sqlmock, ping models.PingResult)
		expectedError error
	}{
		{
			name: "Success",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  now,
					Valid: true,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestAddPingQuery).
					WithArgs(ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Database Error",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  now,
					Valid: true,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestAddPingQuery).
					WithArgs(ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing).
					WillReturnError(errors.New("database error"))
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock, tt.ping)

			err := repo.AddPing(tt.ping)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdatePing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewPostgres(db)

	now := time.Now()
	tests := []struct {
		name          string
		ping          models.PingResult
		mockBehavior  func(mock sqlmock.Sqlmock, ping models.PingResult)
		expectedError error
	}{
		{
			name: "Success with Date",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  now,
					Valid: true,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestUpdatePingWithDateQuery).
					WithArgs(ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedError: nil,
		},
		{
			name: "Success without Date",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Valid: false,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestUpdatePingQuery).
					WithArgs(ping.IPAddress, ping.PingTime).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedError: nil,
		},
		{
			name: "No Rows Affected",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  now,
					Valid: true,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestUpdatePingWithDateQuery).
					WithArgs(ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: sql.ErrNoRows,
		},
		{
			name: "Database Error",
			ping: models.PingResult{
				IPAddress: "192.168.1.1",
				PingTime:  10,
				DateSuccessfulPing: pq.NullTime{
					Time:  now,
					Valid: true,
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock, ping models.PingResult) {
				mock.ExpectExec(TestUpdatePingWithDateQuery).
					WithArgs(ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing).
					WillReturnError(errors.New("database error"))
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock, tt.ping)

			err := repo.UpdatePing(tt.ping)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
