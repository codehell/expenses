package controllers

import (
	"expenses/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"log"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var reg models.Register
	var user models.User
	err := c.BindJSON(&reg)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user.Password = reg.Password
	user.Email = reg.Email
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
