package controllers

import (
	"expenses/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getUserFromClaims(c *gin.Context) (*models.User, bool) {
	email, _ := c.Get("email")

	user := &models.User{
		Email:    email.(*models.User).Email,
		Expenses: []*models.Expense{},
	}
	if err := user.GetUserByEmail(); err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return user, false
	}

	return user, true
}
