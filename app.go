package main

import (
	"expenses/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	e := routes.Router(r)
	_ = e.Run(":8080")
}
