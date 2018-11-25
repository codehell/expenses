package routes

import (
	"github.com/codehell/expenses/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.POST("/auth/register", controllers.RegisterUser)
}