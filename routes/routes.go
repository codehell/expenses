package routes

import (
	"log"
	"time"

	"expenses/controllers"
	"expenses/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
	r.Static("/img", "./public/img")
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})
	corsObject := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposeHeaders:    []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	r.Use(corsObject)
	authMiddleware, err := middlewares.JwtMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.POST("/auth/register", controllers.RegisterUser)
	r.POST("/auth/login", authMiddleware.LoginHandler)
	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		exp := auth.Group("expenses")
		exp.GET("", controllers.ListExpenses)
		exp.POST("", controllers.StoreExpense)
		exp.DELETE("", controllers.DeleteExpense)
		tag := auth.Group("tags")
		tag.POST("", controllers.StoreTag)
	}

	return r
}
