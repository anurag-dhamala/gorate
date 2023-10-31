package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var user = User{
		Name: "Anurag",
	}

	userByte, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)
	w.Write(userByte)
}
