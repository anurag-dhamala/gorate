package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var user = User{
		name: "Anurag",
	}
	userByte, err := json.Marshal(user)
	if err != nil {
		log.Fatal("error occurred")
	}
	w.Write(userByte)
}
