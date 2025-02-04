package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/DurkaVerder/models"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	return db, nil
}

func (p *Postgres) GetAllPing() ([]models.PingResult, error) {
	rows, err := p.db.Query(AllPingQuery)
	if err != nil {
		log.Printf("Error getting all pings: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var pings []models.PingResult
	for rows.Next() {
		var ping models.PingResult
		err = rows.Scan(&ping.IPAddress, &ping.PingTime, &ping.DateSuccessfulPing)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		pings = append(pings, ping)
	}

	return pings, nil
}

func (p *Postgres) AddPing(ping models.PingResult) error {
	_, err := p.db.Exec(AddPingQuery, ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing)
	if err != nil {
		log.Printf("Error adding ping: %v\n", err)
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
