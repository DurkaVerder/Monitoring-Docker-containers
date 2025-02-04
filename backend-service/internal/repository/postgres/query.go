package postgres

const (
	AllPingQuery = "SELECT * FROM ping_result"
	AddPingQuery = "INSERT INTO ping_result (ip_address, ping_time, date_successful_ping) VALUES ($1, $2, $3)"
)
