package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// User
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "OI",
		})
	})
}
