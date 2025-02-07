package postgres

const (
	GetPingQuery            = "SELECT * FROM ping_result WHERE ip_address = $1"
	AllPingQuery            = "SELECT * FROM ping_result"
	AddPingQuery            = "INSERT INTO ping_result (ip_address, ping_time, date_successful_ping) VALUES ($1, $2, $3)"
	UpdatePingWithDateQuery = "UPDATE ping_result SET ping_time = $2, date_successful_ping = $3 WHERE ip_address = $1;"
	UpdatePingQuery         = "UPDATE ping_result SET ping_time = $2 WHERE ip_address = $1;"

	TestGetPingQuery    = "SELECT \\* FROM ping_result WHERE ip_address = \\$1"
	TestAllPingQuery    = "SELECT \\* FROM ping_result"
	TestAddPingQuery    = "INSERT INTO ping_result \\(ip_address, ping_time, date_successful_ping\\) VALUES \\(\\$1, \\$2, \\$3\\)"
	TestUpdatePingQuery = "UPDATE ping_result SET ping_time = \\$2, date_successful_ping = CASE WHEN \\$3 IS NOT NULL THEN \\$3 ELSE date_successful_ping END WHERE ip_address = \\$1;"
)
