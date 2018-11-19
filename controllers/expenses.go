package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/codehell/expenses/models"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func ListExpenses(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	email := fmt.Sprintf("%v", claims["email"])
	user := &models.User{
		Email: email,
	}
	err = user.GetUserByEmail()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	err = user.GetUserExpenses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	user.Password = ""
	userJson, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(userJson)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gasto requerido"))
}