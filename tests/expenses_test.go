package tests

import (
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

func TestUserExpensesList(t *testing.T) {
	email := "admin@codehell.net"
	user := models.User{
		Email:    email,
		Password: "secret",
	}
	db := models.GetDb()
	_, err := db.Model(&user).Insert()
	if err != nil {
		log.Fatalln("Test fail faking db:", err)
	}
	r := getRouter()
	token := makeTokenString("", email)
	req, _ := http.NewRequest("GET", "/expenses", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	clearUserTable(db)
}