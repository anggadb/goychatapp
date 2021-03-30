package controllers

import (
	"database/sql"
	"fmt"
	"goychatapp/lib"
	"goychatapp/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	hashedPassword, err := lib.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user.Password = hashedPassword
	res, err := models.CreateUser(user)
	link := "Please click this link " + os.Getenv("DOMAIN") + "?token=" + res
	mail := lib.SendMail("Verify Account", "Goy System <accdev.bachtiar@gmail.com>", link, []string{user.Email}, []string{}, []string{})
	if mail != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": mail.Error()})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data":    res,
		"message": "Silahkan cek email anda",
	})
}
func GetProfile(c *gin.Context) {
	email := fmt.Sprintf("%v", c.MustGet("email"))
	res, err := models.GetUser(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": res})
}
func UpdateProfile(c *gin.Context) {
	email := fmt.Sprintf("%v", c.MustGet("email"))
	res, err := models.GetUser(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.ShouldBind(&res)
	rows, err := models.UpdateUser(res.Email, res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rows})
}
func DeleteAccount(c *gin.Context) {
	email := fmt.Sprintf("%v", c.MustGet("email"))
	res, err := models.DeleteUser(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if res == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Akun tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
func Login(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	tempPassword := user.Password
	res, err := models.GetUser(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Username tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	userType := res.Type
	if userType != "user" {
		userType = os.Getenv("USER_TOKEN")
	} else {
		userType = os.Getenv("ADMIN_TOKEN")
	}
	match := lib.PasswordMatcher(res.Password, tempPassword)
	if !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password tidak cocok"})
		return
	}
	expiresIn := time.Now().Add(96 * time.Hour)
	claims := lib.Payload{uint(res.ID), res.Email, res.Type, jwt.StandardClaims{ExpiresAt: expiresIn.Unix()}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	encodedToken, err := token.SignedString([]byte(os.Getenv("USER_TOKEN")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	res.Password = ""
	c.JSON(http.StatusOK, gin.H{"token": encodedToken, "data": res})
}
