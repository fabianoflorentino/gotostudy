package routes

import (
	"net/http"

	"github.com/fabianoflorentino/gotostudy/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	routes := []struct {
		route   string
		handler func(w http.ResponseWriter, r *http.Request)
		method  string
	}{
		{"/health", handlers.Health, http.MethodGet},
	}

	for _, rt := range routes {
		r.HandleFunc(rt.route, rt.handler).Methods(rt.method)
	}

	return r
}
