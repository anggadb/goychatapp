package controllers

import (
	"goychatapp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Fatalf("Error while fetch data. %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
