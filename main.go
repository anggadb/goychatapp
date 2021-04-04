package main

import (
	"log"
	"os"

	c "goychatapp/controllers"
	r "goychatapp/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()
	r.UserRouter(router.Group(os.Getenv("API_VERSION")))
	r.AdminRouters(router.Group(os.Getenv("API_VERSION")))
	r.FileRouters(router.Group(os.Getenv("API_VERSION")))
	router.GET("verify-account", c.VerifyAccount)
	router.Static("image", wd+"/assets/images/")
	router.Run(os.Getenv("PORT"))
}
