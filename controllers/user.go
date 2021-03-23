package controllers

import (
	"goychatapp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responses struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Fatalf("Error while fetch data. %v", err)
	}
	c.JSON(http.StatusOK, users)
}
