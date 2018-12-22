package tests

import (
	"bytes"
	"encoding/json"
	"expenses/models"
	"expenses/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	// Having
	r := routes.Router()
	var register models.Register
	register.Email = "cazaplanetas@gmail.com"
	register.Password = "secret"
	register.ConfirmPassword = "secret"
	jsonUser, _ := json.Marshal(&register)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonUser))
	r.ServeHTTP(w, req)

	// Test
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "The user was registered", w.Body.String())

	// Reset
	db := models.GetDb()
	defer db.Close()
	_, err := db.Model((*models.User)(nil)).Where("true").Delete()
	if err != nil {
		log.Fatalln(err)
	}
}
