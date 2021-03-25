package routers

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {
	route.POST("user", c.CreateUser)
	route.GET("users", jwt.AdminAuth, c.GetAllUsers)
	route.POST("user/login", c.Login)
}
