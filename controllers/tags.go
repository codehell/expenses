package controllers

import (
	"github.com/codehell/expenses/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StoreTag(c *gin.Context) {
	var tag models.Tag
	user, ok := getUserFromClaims(c)
	if !ok {
		return
	}
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	tag.UserId = user.Id
	if err := tag.StoreTag(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
