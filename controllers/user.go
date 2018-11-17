package controllers

import (
	"encoding/json"
	"github.com/codehell/expenses/models"
	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error: Can't process data", http.StatusUnprocessableEntity)
		log.Fatal(err)
	}
	err = user.StoreUser()
	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		http.Error(w, "Error: User already exist", http.StatusConflict)
	} else if ok {
		http.Error(w, "Error: Can't register user", http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Error: Can't process data", http.StatusUnprocessableEntity)
	}
	user := models.User{Email: email}
	user.GetUser()
	hashed := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		http.Error(w, "Error: Incorrect user or password", http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusOK)
}