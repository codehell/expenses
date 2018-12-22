package controllers

import (
	"expenses/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StoreTag(c *gin.Context) {
	var tag models.Tag
	user, ok := getUserFromClaims(c)
	if !ok {
		return
	}
	err := c.BindJSON(&tag)
	if err != nil {
		log.Println(err)
		return
	}
	tag.UserId = user.Id
	if err := tag.StoreTag(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
