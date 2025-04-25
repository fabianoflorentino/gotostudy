// Package app provides the core application container for initializing and managing
// dependencies required by the application. It includes the setup of the database
// connection, repository, and services, ensuring that all components are properly
// configured and ready to use.
package app

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/adapters/outbound/persistence/postgres"
	"github.com/fabianoflorentino/gotostudy/core/domain"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/fabianoflorentino/gotostudy/database"
	"gorm.io/gorm"
)

// AppContainer is a struct that serves as a dependency injection container
// for the application. It holds references to shared resources and services
// that are used throughout the application, such as the database connection
// (DB) and the UserService for managing user-related operations.
type AppContainer struct {
	DB          *gorm.DB
	UserService *services.UserService
}

// NewAppContainer initializes and returns a new instance of AppContainer.
// It sets up the database connection, initializes the user repository and service,
// and performs database migrations for the User domain model. If any errors occur
// during database initialization or migration, they are logged. The returned
// AppContainer includes the database connection and the user service.
func NewAppContainer() *AppContainer {
	if err := database.InitDB(); err != nil {
		log.Printf("failed to initialize database: %v", err)
	}

	dbConn := database.DB
	repo := postgres.NewPostgresUserRepository(dbConn)
	service := services.NewUserService(repo)

	if err := dbConn.AutoMigrate(&domain.User{}); err != nil {
		log.Printf("failed to migrate database: %v", err)
	}

	return &AppContainer{
		DB:          dbConn,
		UserService: service,
	}
}
