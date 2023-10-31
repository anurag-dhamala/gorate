package main

import (
	"net/http"

	"github.com/anurag-dhamala/gorate/routes"
)

func main() {
	http.Handle("/", routes.Routes())
	http.ListenAndServe(":8080", nil)
}
