// Package database provides functionality to interact with a PostgreSQL database
// using the GORM library. It includes methods to establish a connection, close
// the connection, and construct the connection string dynamically based on
// environment variables. This package is designed to simplify database operations
// and ensure proper resource management.
package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fabianoflorentino/gotostudy/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	username = os.Getenv("POSTGRES_USER")
	host     = os.Getenv("POSTGRES_HOST")
	password = os.Getenv("POSTGRES_PASSWORD")
	database = os.Getenv("POSTGRES_DB")
	port     = os.Getenv("POSTGRES_PORT")
	sslmode  = os.Getenv("POSTGRES_SSLMODE")
	timezone = os.Getenv("POSTGRES_TIMEZONE")
)

var (
	exists bool
	DB     *gorm.DB
)

// InitDB initializes the database connection using GORM and PostgreSQL.
// It reads the connection parameters from environment variables and sets up
// the database connection. It also enables the pgcrypto extension if it is not
// already enabled. The function logs fatal errors if the connection fails or
// if the extension cannot be enabled.
func InitDB() {
	dsn := setPostgresConnectionString()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Enable the pgcrypto extension
	if err := EnablePgcryptoExtension(db); err != nil {
		log.Fatalf("failed to enable pgcrypto extension: %v", err)
	}

	DB = db
	if err := runMigrations(db,
		models.User{},
		models.Task{},
	); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}

// enablePgcryptoExtension checks if the pgcrypto extension exists and creates it if not.
func EnablePgcryptoExtension(db *gorm.DB) error {
	createQuery := "CREATE EXTENSION IF NOT EXISTS pgcrypto;"
	checkQuery := "SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pgcrypto');"

	// Check if the extension already exists
	if err := db.Exec(createQuery).Error; err != nil {
		return err
	}
	// Check if the extension was created successfully
	if err := db.Raw(checkQuery).Scan(&exists).Error; err != nil {
		return err
	}
	// If the extension was not created, return an error
	if !exists {
		return errors.New("pgcrypto extension not found after creation")
	}

	return nil
}

// setPostgresConnectionString constructs the connection string for PostgreSQL
// using the environment variables defined above.
// It returns a string that can be used to connect to the PostgreSQL database.
func setPostgresConnectionString() string {
	return "user=" + username + " password=" + password + " host=" + host +
		" port=" + port + " dbname=" + database +
		" sslmode=" + sslmode + " TimeZone=" + timezone
}

// runMigrations applies database migrations for the provided models using GORM's AutoMigrate method.
// It iterates over the given models and attempts to migrate each one. If any migration fails,
// it returns an error indicating the model that failed and the reason.
//
// Parameters:
//   - db: A pointer to a gorm.DB instance representing the database connection.
//   - models: A variadic parameter of models (of any type) to be migrated.
//
// Returns:
//   - error: An error if any migration fails, or nil if all migrations succeed.
func runMigrations(db *gorm.DB, models ...any) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate model %T: %v", model, err)
		}
	}

	return nil
}
