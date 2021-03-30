package routers

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {
	route.POST("user", c.CreateUser)
	route.GET("user", jwt.UserAuth, c.GetProfile)
	route.PUT("user/update", jwt.UserAuth, c.UpdateProfile)
	route.POST("user/login", c.Login)
	route.DELETE("user/delete", jwt.UserAuth, c.DeleteAccount)
}
