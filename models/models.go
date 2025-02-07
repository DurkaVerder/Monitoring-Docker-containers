package models

import "database/sql"

type PingResult struct {
	ID                 int          `json:"id"`
	IPAddress          string       `json:"ip_address"`
	PingTime           int          `json:"ping_time"`
	DateSuccessfulPing sql.NullTime `json:"date_successful_ping"`
}
