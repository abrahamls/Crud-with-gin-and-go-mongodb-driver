package main

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// User
	router.POST("/users", controllers.GetUsers)
}
