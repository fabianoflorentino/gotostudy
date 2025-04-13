// Package database provides functionality to interact with a PostgreSQL database
// using the GORM library. It includes methods to establish a connection, close
// the connection, and construct the connection string dynamically based on
// environment variables. This package is designed to simplify database operations
// and ensure proper resource management.
package database

import (
	"errors"
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
	db.AutoMigrate(
		&models.User{},
	)
}

// enablePgcryptoExtension checks if the pgcrypto extension exists and creates it if not
func EnablePgcryptoExtension(db *gorm.DB) error {
	var (
		createQuery = "CREATE EXTENSION IF NOT EXISTS pgcrypto;"
		checkQuery  = "SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pgcrypto');"
	)

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
