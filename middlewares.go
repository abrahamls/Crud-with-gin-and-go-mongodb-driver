package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetMiddlewares(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins: true,
	}))
}
