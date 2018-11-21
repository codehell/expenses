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

func StoreExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	user, err := getUserFromClaims(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	expense.UserId = user.Id
	err = expense.StoreExpense()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getUserFromClaims(r *http.Request) (models.User, error) {
	var user models.User
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return user, err
	}
	email := fmt.Sprintf("%v", claims["email"])
	user.Email = email
	err = user.GetUserByEmail()
	return user, nil
}