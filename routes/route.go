package routes

import (
	"github.com/anurag-dhamala/gorate/handlers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", handlers.HomeHandler).Methods("GET")
	return r
}
