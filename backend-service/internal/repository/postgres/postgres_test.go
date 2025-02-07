package postgres

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/DurkaVerder/models"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddPing(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	p := NewPostgres(db)
// 	date := time.Now()

// 	tests := []struct {
// 		name    string
// 		ping    models.PingResult
// 		wantErr bool
// 	}{
// 		{
// 			name: "Successful ping with all fields",
// 			ping: models.PingResult{
// 				IPAddress:          "192.12.32.1",
// 				PingTime:           10,
// 				DateSuccessfulPing: &date,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Successful ping with missing DateSuccessfulPing",
// 			ping: models.PingResult{
// 				IPAddress: "192.13.43.2",
// 				PingTime:  9999,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Successful ping with missing PingTime and DateSuccessfulPing",
// 			ping: models.PingResult{
// 				IPAddress: "192.13.34.1",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Failed ping due to database error",
// 			ping: models.PingResult{
// 				IPAddress: "192.13.34.1",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.wantErr {
// 				mock.ExpectExec(TestAddPingQuery).
// 					WithArgs(tt.ping.IPAddress, tt.ping.PingTime, tt.ping.DateSuccessfulPing).
// 					WillReturnError(errors.New("error connecting to database"))
// 			} else {
// 				mock.ExpectExec(TestAddPingQuery).
// 					WithArgs(tt.ping.IPAddress, tt.ping.PingTime, tt.ping.DateSuccessfulPing).
// 					WillReturnResult(sqlmock.NewResult(1, 1))
// 			}

// 			err := p.AddPing(tt.ping)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}

// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestGetAllPing(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	p := NewPostgres(db)
// 	date := time.Now()
// 	date2 := date.Add(-time.Hour)

// 	tests := []struct {
// 		name    string
// 		rows    *sqlmock.Rows
// 		want    []models.PingResult
// 		wantErr bool
// 	}{
// 		{
// 			name: "Successful ping with all fields",
// 			rows: sqlmock.NewRows([]string{"ip_address", "ping_time", "date_successful_ping"}).
// 				AddRow("192.168.0.1", 23, &date).
// 				AddRow("192.168.0.2", 45, date2),
// 			want: []models.PingResult{
// 				{
// 					IPAddress:          "192.168.0.1",
// 					PingTime:           23,
// 					DateSuccessfulPing: &date,
// 				},
// 				{
// 					IPAddress:          "192.168.0.2",
// 					PingTime:           45,
// 					DateSuccessfulPing: &date2,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Empty result set",
// 			rows:    sqlmock.NewRows([]string{"ip_address", "ping_time", "date_successful_ping"}),
// 			want:    []models.PingResult(nil),
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Database error",
// 			rows:    nil,
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.wantErr {
// 				mock.ExpectQuery(TestAllPingQuery).
// 					WillReturnError(errors.New("error connecting to database"))
// 			} else {
// 				mock.ExpectQuery(TestAllPingQuery).
// 					WillReturnRows(tt.rows)
// 			}

// 			pings, err := p.GetAllPing()

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Nil(t, pings)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.want, pings)
// 			}

// 		})
// 	}
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestGetPing(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	p := NewPostgres(db)
// 	date := time.Now()

// 	tests := []struct {
// 		name    string
// 		ip      string
// 		row     *sqlmock.Rows
// 		want    *models.PingResult
// 		wantErr bool
// 	}{
// 		{
// 			name: "Successful ping with all fields",
// 			ip:   "192.168.0.1",
// 			row: sqlmock.NewRows([]string{"ip_address", "ping_time", "date_successful_ping"}).
// 				AddRow("192.168.0.1", 23, &date),
// 			want: &models.PingResult{
// 				IPAddress:          "192.168.0.1",
// 				PingTime:           23,
// 				DateSuccessfulPing: &date,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Empty result set",
// 			ip:      "192.168.0.1",
// 			row:     sqlmock.NewRows([]string{"ip_address", "ping_time", "date_successful_ping"}),
// 			want:    nil,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Database error",
// 			ip:      "192.168.0.1",
// 			row:     nil,
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.wantErr {
// 				mock.ExpectQuery(TestGetPingQuery).
// 					WithArgs(tt.ip).
// 					WillReturnError(errors.New("error connecting to database"))
// 			} else {
// 				mock.ExpectQuery(TestGetPingQuery).
// 					WithArgs(tt.ip).
// 					WillReturnRows(tt.row)
// 			}

// 			ping, err := p.GetPing(tt.ip)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Nil(t, ping, "Expected nil result when error occurs")
// 			} else {
// 				assert.NoError(t, err)

// 				if tt.want == nil {
// 					assert.Nil(t, ping, "Expected nil result for empty rows")
// 				} else {
// 					assert.Equal(t, tt.want, ping, "Result mismatch")
// 				}
// 			}

// 		})
// 	}
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestUpdatePing(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	p := NewPostgres(db)
// 	date := time.Now()

// 	tests := []struct {
// 		name    string
// 		ping    models.PingResult
// 		wantErr bool
// 	}{
// 		{
// 			name: "Successful ping with all fields",

// 			ping: models.PingResult{
// 				IPAddress:          "192.168.0.1",
// 				PingTime:           23,
// 				DateSuccessfulPing: &date,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Empty result set",
// 			ping: models.PingResult{
// 				IPAddress: "192.168.0.1",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Database error",
// 			ping: models.PingResult{
// 				IPAddress: "192.168.0.1",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.wantErr {
// 				mock.ExpectExec(TestUpdatePingQuery).
// 					WithArgs(tt.ping.IPAddress, tt.ping.PingTime, tt.ping.DateSuccessfulPing).
// 					WillReturnError(errors.New("error connecting to database"))
// 			} else {
// 				mock.ExpectExec(TestUpdatePingQuery).
// 					WithArgs(tt.ping.IPAddress, tt.ping.PingTime, tt.ping.DateSuccessfulPing).
// 					WillReturnResult(sqlmock.NewResult(0, 1))
// 			}

// 			err := p.UpdatePing(tt.ping)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}

// 	assert.NoError(t, mock.ExpectationsWereMet())
// }
