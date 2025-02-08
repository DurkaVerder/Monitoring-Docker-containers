package postgres

const (
	GetPingQuery            = "SELECT * FROM ping_result WHERE ip_address = $1"
	AllPingQuery            = "SELECT * FROM ping_result"
	AddPingQuery            = "INSERT INTO ping_result (ip_address, ping_time, date_successful_ping) VALUES ($1, $2, $3)"
	UpdatePingWithDateQuery = "UPDATE ping_result SET ping_time = $2, date_successful_ping = $3 WHERE ip_address = $1;"
	UpdatePingQuery         = "UPDATE ping_result SET ping_time = $2 WHERE ip_address = $1;"
)
