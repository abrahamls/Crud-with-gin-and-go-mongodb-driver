package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// load the env here before calling
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := os.Getenv("SERVER_ADDRESS")

	router := gin.Default()

	SetMiddlewares(router)
	SetRoutes(router)

	logrus.Error(router.Run(server))
}
