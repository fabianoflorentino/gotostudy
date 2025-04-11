package migration

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/model"
)

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

	database.EnablePgcryptoExtension(conn)

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
