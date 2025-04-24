package app

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/controllers"
	"github.com/fabianoflorentino/gotostudy/adapters/outbound/persistence/postgres"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/fabianoflorentino/gotostudy/database"
	"gorm.io/gorm"
)

type AppContainer struct {
	DB          *gorm.DB
	UserService *services.UserService
}

func New() *AppContainer {
	if err := database.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	dbConn := database.DB
	repo := postgres.NewPostgresUserRepository(dbConn)
	service := services.NewUserService(repo)

	if err := dbConn.AutoMigrate(&controllers.UserController{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return &AppContainer{
		DB:          dbConn,
		UserService: service,
	}
}
