package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Up!"))
	w.WriteHeader(http.StatusOK)
}
