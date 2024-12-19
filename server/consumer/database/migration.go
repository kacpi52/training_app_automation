package consumer_database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func connectWithRetry(db *sql.DB, maxRetries int, delay time.Duration) error {
	var err error

	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}

		log.Printf("Database connection failed: %v. Retrying in %s...", err, delay)
		time.Sleep(delay)
	}

	return fmt.Errorf("failed to verify database connection after %d attempts: %w", maxRetries, err)
}

func RunMigration(db *sql.DB) error {
	var tableExists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'post'
		)
	`

	err := connectWithRetry(db, 10, 5*time.Second)
	if err != nil {
		return fmt.Errorf("unable to verify database connection: %w", err)
	}

	err = db.QueryRow(query).Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("failed to check table existence: %w", err)
	}

	if tableExists {
		log.Println("Table 'post' already exists. Skipping migrations.")
		return nil
	}
	// Konfiguracja migracji
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://consumer/database/migration",
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Wykonywanie migracji
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations ran successfully!")
	return nil
}