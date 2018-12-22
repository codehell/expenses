package tests

import (
	"bytes"
	"encoding/json"
	"expenses/models"
	"expenses/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	r := routes.Router(e)
	// Having
	var register models.Register
	register.Email = "admin@codehell.net"
	register.Password = "secret"
	register.ConfirmPassword = "secret"
	jsonUser, _ := json.Marshal(&register)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user models.User
	db := models.GetDb()
	_ = db.Model(&user).Where("email = ?", "admin@codehell.net").Select()

	// Tests
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "The user was registered", w.Body.String())
	assert.Equal(t, "admin@codehell.net", user.Email)

	ClearUserTable(db)
}

func TestLoginUser (t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	r := routes.Router(e)

	user := models.User{
		Email: "admin@codehell.net",
		Password: "secret",
	}
	db := models.GetDb()

	_, err := db.Model(&user).Insert()
	if err != nil {
		log.Fatalln("Test fail faking db:", err)
	}
	req,  _ := http.NewRequest("POST", "/auth/login", nil)
	w := httptest.NewRecorder()
	req.SetBasicAuth("admin@codehell.net", "secret")
	r.ServeHTTP(w, req)

	// Tests
	assert.Equal(t, http.StatusOK, w.Code)

	ClearUserTable(db)
}

func ClearUserTable (db *pg.DB) {
	_, err := db.Model((*models.User)(nil)).Where("true").Delete()
	if err != nil {
		log.Fatalln(err)
	}
}
