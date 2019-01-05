package tests

import (
	"bytes"
	"encoding/json"
	"expenses/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	r := getRouter()
	// Having
	var register models.Register
	register.Email = "admin@codehell.net"
	register.Password = "secret"
	register.ConfirmPassword = "secret"
	jsonUser, _ := json.Marshal(&register)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user models.User
	db := models.GetDb()
	_ = db.Model(&user).Where("email = ?", "admin@codehell.net").Select()

	// Tests
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "The user was registered", w.Body.String())
	assert.Equal(t, "admin@codehell.net", user.Email)

	clearUserTable(db)
}

func TestLoginUser(t *testing.T) {

	db := models.GetDb()
	createTestUser(db)
	defer clearUserTable(db)

	r := getRouter()
	req, _ := http.NewRequest("POST", "/auth/login", nil)
	req.SetBasicAuth("admin@codehell.net", "secret")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Tests
	assert.Equal(t, http.StatusOK, w.Code)
}
