package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/codehell/expenses/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
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
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Error: Can't process data", http.StatusUnprocessableEntity)
		return
	}
	user := models.User{Email: email}
	user.GetUserByEmail()
	hashed := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		http.Error(w, "Error: Incorrect user or password", http.StatusUnauthorized)
		return
	}
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	mapClaims := jwt.MapClaims{
		"email": user.Email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}
	_, tokenString, err := tokenAuth.Encode(jwtauth.Claims(mapClaims))
	if err != nil {
		http.Error(w, "Error: Incorrect user or password", http.StatusInternalServerError)
		log.Fatal("Error: Can't generate token")
		return
	}
	fmt.Fprint(w, tokenString)
}