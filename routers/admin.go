package routers

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func AdminRouters(route *gin.RouterGroup) {
	route.GET("users", jwt.AdminAuth, c.GetAllUsers)
}
