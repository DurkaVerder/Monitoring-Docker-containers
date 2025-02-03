package models

import "time"

type PingResult struct {
	ID                 int       `json:"id" db:"id"`
	IPAddress          string    `json:"ip_address" db:"ip_address"`
	PingTime           int       `json:"ping_time" db:"ping_time"`
	DateSuccessfulPing time.Time `json:"date_successful_ping" db:"date_successful_ping"`
}
