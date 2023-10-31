package main

import (
	"net/http"
	"time"

	"github.com/anurag-dhamala/gorate/routes"
)

func main() {

	srv := &http.Server{
		Handler: routes.Routes(),
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
