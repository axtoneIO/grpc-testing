package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Migrate - will execute the migration task according to the files contained
// on the migrations directory
func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil{
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil{
		return err
	}

	if err := m.Up(); err != nil{
		if err.Error() == "no change"{
			log.Println("no change made by migrations") 
		} else {
			return err
		}
	}

	return nil
}
