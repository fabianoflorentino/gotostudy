package main

import (
	"github.com/fabianoflorentino/gotostudy/config"
	"github.com/fabianoflorentino/gotostudy/database"
	"github.com/fabianoflorentino/gotostudy/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.InitDB()
}

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)

	r.Run()
}
