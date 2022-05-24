package db

import (
	"fmt"
	"os"

	"github.com/axtoneIO/grpc-testing/internal/rocket"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

// New - returns a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

// GetRocket - retrieves a rocket from the database by id (db implementation)
func (s Store) GetRocket(id int64) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(
		`SELECT id, name, type FROM rockets WHERE id =$1;`,
		id,
	)
	err := row.Scan(
		&rkt.Id,
		&rkt.Name,
		&rkt.Type,
	)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

// AddRocket - Inserts a rocket into the rockets table
func (s Store) AddRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets
		(id, name, type)
		VALUES (:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		return rocket.Rocket{}, err
	}
	return rocket.Rocket{
		Id:   rkt.Id,
		Name: rkt.Name,
		Type: rkt.Type,
	}, nil
}

// DeleteRocket - Deletes a rocket from the rockets table
func (s Store) DeleteRocket(id int64) (string, error) {
	_, err := s.db.Exec(`DELETE FROM rockets WHERE id=$1`, id)
	if err != nil {
		return "", err
	}
	return "Rocket deleted successfully", nil
}
