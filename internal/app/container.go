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
	TaskService *services.TaskService
}

// NewAppContainer initializes and returns a new instance of AppContainer.
// It sets up the database connection, initializes the user repository and service,
// and performs database migrations for the User domain model. If any errors occur
// during database initialization or migration, they are logged. The returned
// AppContainer includes the database connection and the user service.
func NewAppContainer() *AppContainer {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("failed to initialize database: %v", err)
		return nil
	}

	usrService := usrService(db)
	tskService := tskService(db)

	return &AppContainer{
		DB:          db,
		UserService: usrService,
		TaskService: tskService,
	}
}

func usrService(db *gorm.DB) *services.UserService {
	usr := postgres.NewPostgresUserRepository(db)
	srv := services.NewUserService(usr)

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Printf("failed to migrate user repository: %v", err)
	}

	return srv
}

func tskService(db *gorm.DB) *services.TaskService {
	tsk := postgres.NewPostgresTaskRepository(db)
	usr := postgres.NewPostgresUserRepository(db)
	tskService := services.NewTaskService(tsk, usr)

	if err := db.AutoMigrate(&domain.Task{}); err != nil {
		log.Printf("failed to migrate task repository: %v", err)
	}

	return tskService
}
