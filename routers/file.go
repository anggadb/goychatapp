package routers

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"

	"github.com/gin-gonic/gin"
)

func FileRouters(route *gin.RouterGroup) {
	route.POST("file", jwt.UserAdminAuth, c.UploadFile)
	route.GET("files", jwt.UserAuth, c.GetAllFiles)
	route.DELETE("file/:id/delete", jwt.UserAdminAuth, c.DeleteFile)
}
