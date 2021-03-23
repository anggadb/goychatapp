package routers

import (
	c "goychatapp/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {
	route.POST("user", c.CreateUser)
	route.GET("users", c.GetAllUsers)
}
