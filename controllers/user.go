package controllers

import (
	"expenses/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
)

type register struct {
	Email string `json:"email" binding:"required,email,max=124"`
	Password string `json:"password" binding:"required,max=124"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=ConfirmPassword,max=124"`
}

func RegisterUser(c *gin.Context) {
	var reg register
	var user models.User
	err := c.BindJSON(&reg)
	if err != nil {
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
