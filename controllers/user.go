package controllers

import (
	"encoding/json"
	"expenses/models"
	"github.com/go-pg/pg"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Can't process data", http.StatusUnprocessableEntity)
	}
	err = user.StoreUser()
	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		http.Error(w, "User already exist", http.StatusConflict)
	}
	w.WriteHeader(http.StatusCreated)
}
