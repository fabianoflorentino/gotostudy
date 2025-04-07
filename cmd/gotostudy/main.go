package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fabianoflorentino/gotostudy/database/migration"
)

func init() {
	migration.Run()
}

func main() {
	log.Fatal(http.ListenAndServe(os.Getenv("GTS_LOCAL_PORT"), nil).Error())
}
