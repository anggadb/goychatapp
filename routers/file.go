package routers

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func FileRouters(route *gin.RouterGroup) {
	route.POST("file", jwt.UserAuth, c.UploadFile)
}
