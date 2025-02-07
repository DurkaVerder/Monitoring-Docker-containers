package models

import "github.com/lib/pq"

type PingResult struct {
	ID                 int         `json:"id"`
	IPAddress          string      `json:"ip_address"`
	PingTime           int         `json:"ping_time"`
	DateSuccessfulPing pq.NullTime `json:"date_successful_ping"`
}
