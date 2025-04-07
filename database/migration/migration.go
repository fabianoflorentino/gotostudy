package migration

import (
	"errors"
	"log"

	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/model"
	"gorm.io/gorm"
)

const (
	createQuery = "CREATE EXTENSION IF NOT EXISTS pgcrypto;"
	checkQuery  = "SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pgcrypto');"
)

var exists bool

// Run performs the database migration for the application using GORM
func Run() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("failed to create database instance: %v", err)
		return
	}

	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	enablePgcryptoExtension(conn)

	// Perform GORM auto-migration
	err = conn.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatalf("failed to perform auto-migration: %v", err)
		return
	}

	log.Println("successfully performed auto-migration")
}

// enablePgcryptoExtension checks if the pgcrypto extension exists and creates it if not
func enablePgcryptoExtension(db *gorm.DB) error {

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
