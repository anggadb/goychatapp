package controllers

import (
	"database/sql"
	"fmt"
	"goychatapp/models"
	"net/http"
	"strconv"

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
func GetAllRooms(c *gin.Context) {
	var filters models.Room
	c.Bind(&filters)
	userId := c.MustGet("id").(uint)
	filters.SenderId = userId
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	rooms, err := models.GetAllRooms(filters, c.Query("order_by"), c.Query("order"), page, perPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms, "page": page, "per_page": perPage, "count": len(rooms)})
}
