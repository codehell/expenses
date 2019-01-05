package tests

import (
	"encoding/json"
	"expenses/models"
	"expenses/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusUnauthorizedWithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	r := routes.Router(e)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/expenses", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func StoreExpense(t *testing.T) {
	db := models.GetDb()
	user := createTestUser(db)
	expense := models.Expense{
		UserId: user.Id,
		Amount: "10.25",
		Description: "a expense test",
	}
	expenseJ, _ := json.Marshal(expense)
	log.Println(expenseJ)
}

func TestUserExpensesList(t *testing.T) {
	db := models.GetDb()
	email := "admin@codehell.net"
	user := models.User{
		Email:    email,
		Password: "secret",
	}
	_, err := db.Model(&user).Insert()
	if err != nil {
		log.Fatalln("Test fail faking user: ", err)
	}
	expenses := []models.Expense{
		models.Expense{
			UserId: user.Id,
			Amount: "10.25",
			Description: "a expense test",
		},
		models.Expense{
			UserId: user.Id,
			Amount: "100000",
			Description: "a second expense test",
		},
	}
	_, err = db.Model(&expenses).Insert()
	if err != nil {
		log.Fatalln("Test fail faking expenses: ", err)
	}

	w, _ := callRouteWithToken("/expenses", "GET", user.Email)

	// Get created expense
	_ = db.Model(&expenses).Select()

	// Clear database
	defer clearUserTable(db)
	defer clearExpensesTable(db)

	// Get response and turn json
	var u models.User
	err = json.NewDecoder(w.Body).Decode(&u)
	if err != nil {
		log.Fatalln("Error in body response")
	}

	// Make assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, u.Email, user.Email)
	assert.Equal(t, *u.Expenses[0], expenses[0])
	assert.Equal(t, *u.Expenses[1], expenses[1])
}
