package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fabianoflorentino/gotostudy/database/migration"
	"github.com/fabianoflorentino/gotostudy/routes"
)

func init() {
	if os.Getenv("GTS_LOCAL_PORT") == "" {
		log.Fatal("GTS_LOCAL_PORT environment variable is not set")
	}

	migration.Run()
}

func main() {
	r := routes.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("GTS_LOCAL_PORT"), r))
}
