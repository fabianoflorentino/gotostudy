package app

import (
	"log"

	"github.com/fabianoflorentino/gotostudy/adapters/inbound/http/controllers"
	"github.com/fabianoflorentino/gotostudy/adapters/outbound/persistence"
	"github.com/fabianoflorentino/gotostudy/core/services"
	"github.com/fabianoflorentino/gotostudy/database"
	"gorm.io/gorm"
)

type AppContainer struct {
	DB         *gorm.DB
	Controller *controllers.UserController
}

func New() *AppContainer {
	if err := database.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	dbConn := database.DB
	repo := persistence.New(dbConn)
	service := services.New(repo)
	controller := controllers.New(service)

	return &AppContainer{
		DB:         dbConn,
		Controller: controller,
	}
}
