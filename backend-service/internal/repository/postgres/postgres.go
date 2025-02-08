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

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
		return nil
	}

	return db
}

func (p *Postgres) GetPing(IPAdress string) (*models.PingResult, error) {
	var ping models.PingResult
	err := p.db.QueryRow(GetPingQuery, IPAdress).Scan(&ping.ID, &ping.IPAddress, &ping.PingTime, &ping.DateSuccessfulPing)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No ping result found for IP: %s\n", IPAdress)
			return nil, sql.ErrNoRows
		}
		log.Printf("Error getting ping: %v\n", err)
		return nil, err
	}

	return &ping, nil
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
		err = rows.Scan(&ping.ID, &ping.IPAddress, &ping.PingTime, &ping.DateSuccessfulPing)
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

func (p *Postgres) UpdatePing(ping models.PingResult) error {
	var result sql.Result
	var err error
	if ping.DateSuccessfulPing.Valid {
		result, err = p.db.Exec(UpdatePingWithDateQuery, ping.IPAddress, ping.PingTime, ping.DateSuccessfulPing)
	} else {
		result, err = p.db.Exec(UpdatePingQuery, ping.IPAddress, ping.PingTime)
	}
	if err != nil {
		log.Printf("Error updating ping: %v\n", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v\n", err)
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
