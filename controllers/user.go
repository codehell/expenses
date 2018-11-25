package controllers

import (
	"fmt"
	"github.com/codehell/expenses/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	err = user.StoreUser()
	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		c.JSON(http.StatusConflict, gin.H{
			"status": http.StatusConflict,
			"error": "User already exist",
		})
	} else if ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error": "Can't register user",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": http.StatusCreated,
			"message": "The user was registered",
		})
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