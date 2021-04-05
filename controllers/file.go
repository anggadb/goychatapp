package controllers

import (
	"fmt"
	"goychatapp/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	os, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	userID := c.MustGet("id").(uint)
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext
	err = c.SaveUploadedFile(file, os+"/assets/images/"+fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	files := models.Files{
		UserId: userID,
		Name:   fileName,
		Type:   "user-photo",
		Path:   "/image/" + fileName,
	}
	path, err := models.CreateFile(files)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": path})
}

func GetAllFiles(c *gin.Context) {
	var file models.Files
	c.Bind(&file)
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
	files, err := models.GetAllFiles(file, c.Query("order_by"), c.Query("order"), page, perPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": files, "page": page, "per_page": perPage, "count": len(files)})
}
func DeleteFile(c *gin.Context) {
	id := c.Params.ByName("id")
	userID := c.MustGet("id").(uint)
	userType := fmt.Sprintf("%v", c.MustGet("type"))
	if userType == "user" {
		res, err := models.GetFile(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if res.UserId != userID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "File anda tidak ditemukan"})
			return
		}
	}
	err := models.DeleteFile(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus file"})
}
