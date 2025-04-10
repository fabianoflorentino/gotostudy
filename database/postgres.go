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

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	createQuery = "CREATE EXTENSION IF NOT EXISTS pgcrypto;"
	checkQuery  = "SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pgcrypto');"
)

var exists bool

var (
	username = os.Getenv("POSTGRES_USER")
	host     = os.Getenv("POSTGRES_HOST")
	password = os.Getenv("POSTGRES_PASSWORD")
	database = os.Getenv("POSTGRES_DB")
	port     = os.Getenv("POSTGRES_PORT")
	sslmode  = os.Getenv("POSTGRES_SSLMODE")
	timezone = os.Getenv("POSTGRES_TIMEZONE")
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres() (Database, error) {
	return &Postgres{}, nil
}

// Conn establishes a connection to the PostgreSQL database using GORM.
// It uses the connection string constructed by setPostgresConnectionString.
// If the connection is successful, it returns a pointer to the gorm.DB object.
func (p *Postgres) Connect() (*gorm.DB, error) {
	if p.db != nil {
		return p.db, nil
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  setPostgresConnectionString(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Database connection established")
	p.db = db

	return db, nil
}

// Close closes the database connection.
// It retrieves the database connection using the Conn method and then
// calls the Close method on the underlying sql.DB object.
func (p *Postgres) Close() error {
	if p.db == nil {
		return nil
	}

	sqlDB, err := p.db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB: %v", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("failed to close sqlDB: %v", err)
		return err
	}
	defer sqlDB.Close()

	log.Println("Database connection closed")
	return nil
}

// enablePgcryptoExtension checks if the pgcrypto extension exists and creates it if not
func EnablePgcryptoExtension(db *gorm.DB) error {

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
