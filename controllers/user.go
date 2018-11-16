package controllers

import (
	"encoding/json"
	"invernalia/models"
	"log"
	"net/http"
)

func RegisterUser (w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Can't process data", http.StatusUnprocessableEntity)
	}
	err = user.StoreUser()
	if err != nil {
		log.Fatal(err)
	}
}