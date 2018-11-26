package controllers

import (
	"github.com/codehell/expenses/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	err = user.StoreUser()
	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		c.String(http.StatusConflict, "User already exist")
	} else if ok {
		c.String(http.StatusInternalServerError, "Can't register user")
	} else {
		c.String(http.StatusCreated, "The user was registered")
	}
}
