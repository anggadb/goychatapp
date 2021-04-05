package lib

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Payload struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.StandardClaims
}

func AdminAuth(c *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	token := c.Request.Header.Get("Authorization")
	claims := &Payload{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ADMIN_TOKEN")), nil
	})
	if err != nil {
		if err.Error() == "signature is invalid" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Gagal memverifikasi algoritma token",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
		return
	}
	c.Set("id", claims.ID)
	c.Set("type", claims.Type)
	c.Set("email", claims.Email)
	c.Next()
}
func UserAuth(c *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	token := c.Request.Header.Get("Authorization")
	claims := &Payload{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("USER_TOKEN")), nil
	})
	if err != nil {
		if err.Error() == "signature is invalid" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Gagal memverifikasi algoritma token",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
		return
	}
	c.Set("id", claims.ID)
	c.Set("type", claims.Type)
	c.Set("email", claims.Email)
	c.Next()
}
func UserAdminAuth(c *gin.Context) {
	var parsedToken *jwt.Token
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	token := c.Request.Header.Get("Authorization")
	claims := &Payload{}
	parsedToken, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ADMIN_TOKEN")), nil
	})
	if err != nil {
		if err.Error() == "signature is invalid" {
			parsedToken, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("USER_TOKEN")), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": "Gagal memverifikasi algoritma token",
					})
					return
				}
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
		return
	}
	c.Set("id", claims.ID)
	c.Set("type", claims.Type)
	c.Set("email", claims.Email)
	c.Next()
}
