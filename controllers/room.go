package controllers

import (
	"database/sql"
	"fmt"
	"goychatapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var room models.Room
	c.ShouldBind(&room)
	id := c.MustGet("id").(uint)
	room.SenderId = id
	room.Participants = append(room.Participants, fmt.Sprint(id))
	res, err := models.CreateRoom(room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
func GetRoomById(c *gin.Context) {
	id := c.Params.ByName("id")
	userId := c.MustGet("id").(uint)
	room, err := models.GetRoom(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Room tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	if room.SenderId != userId {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Room tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}
