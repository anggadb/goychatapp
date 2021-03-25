package controllers

import (
	"database/sql"
	"goychatapp/lib"
	"goychatapp/models"
	"log"
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
	res := models.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Fatalf("Error while fetch data. %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
func Login(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	tempPassword := user.Password
	res, err := models.GetUser(user)
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
	claims := lib.Payload{uint(res.ID), res.Username, res.Type, jwt.StandardClaims{ExpiresAt: expiresIn.Unix()}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	encodedToken, err := token.SignedString([]byte(userType))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	res.Password = ""
	c.JSON(http.StatusOK, gin.H{"token": encodedToken, "data": res})
}
