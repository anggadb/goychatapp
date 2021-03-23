package main

import (
	"log"
	"os"

	r "goychatapp/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	router := gin.Default()
	r.UserRouter(router.Group(os.Getenv("API_VERSION")))
	router.Run(os.Getenv("PORT"))
}
