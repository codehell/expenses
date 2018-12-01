package controllers

import (
	"github.com/codehell/expenses/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StoreExpense(c *gin.Context) {
	var expense models.Expense
	user, ok := getUserFromClaims(c)
	if !ok {
		return
	}
	err := c.ShouldBindJSON(&expense)

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	expense.UserId = user.Id
	err = expense.StoreExpense()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func DeleteExpense(c *gin.Context) {
	user, ok := getUserFromClaims(c)
	if !ok {
		return
	}
	err := user.GetUserByEmail()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	var expense models.Expense
	err = c.ShouldBindJSON(&expense)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	err = expense.DeleteExpense()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func ListExpenses(c *gin.Context) {
	user, ok := getUserFromClaims(c)
	if !ok {
		return
	}
	err := user.GetUserByEmail()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	err = user.GetUserExpenses()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
