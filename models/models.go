package models

import "time"

type PingResult struct {
	ID                 int       `json:"id"`
	IPAddress          string    `json:"ip_address"`
	PingTime           int       `json:"ping_time"`
	DateSuccessfulPing time.Time `json:"date_successful_ping"`
}