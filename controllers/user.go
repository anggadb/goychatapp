package controllers

import (
	"goychatapp/lib"
	"goychatapp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	salt := lib.GenerateSalt(16)
	var user models.User
	c.ShouldBind(&user)
	user.Password = lib.HashPassword(user.Password, salt)
	res := models.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Fatalf("Error while fetch data. %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
