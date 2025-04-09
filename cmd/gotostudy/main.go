package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fabianoflorentino/gotostudy/database/migration"
	"github.com/fabianoflorentino/gotostudy/internal/errormsg"
	"github.com/fabianoflorentino/gotostudy/routes"
)

var (
	GTS_LOCAL_PORT string = os.Getenv("GTS_LOCAL_PORT")
)

func init() {
	if err := os.Getenv(GTS_LOCAL_PORT); err == "" {
		log.Fatal(errormsg.ErrEnvNotSet, GTS_LOCAL_PORT)
	}

	migration.Run()
}

func main() {
	r := routes.InitializeRoutes()

	log.Fatal(http.ListenAndServe(":"+os.Getenv(GTS_LOCAL_PORT), r))
}
