package routers

import (
	c "goychatapp/controllers"
	l "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func RoomRouters(route *gin.RouterGroup) {
	route.POST("room", l.UserAuth, c.CreateRoom)
	route.GET("room/:id", l.UserAuth, c.GetRoomById)
	route.GET("rooms", l.UserAuth, c.GetAllRooms)
}
