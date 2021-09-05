package test

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steinfletcher/apitest"
)

func TestGetUser(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load the .env file")
	}
	r := gin.Default()
	r.GET("/v1/user", jwt.UserAdminAuth, c.GetProfile)
	testRoute := &TestApplication{Router: r}
	apitest.New().
		Handler(testRoute.Router).
		Get("/v1/user").
		Headers(map[string]string{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTQsImVtYWlsIjoiYmFjaHRpYXIuYW5nZ2FAZ21haWwuY29tIiwidHlwZSI6InVzZXIiLCJleHAiOjE2MzEyMTQ2NjR9.2LUNgvoxNoEVLjBEQLYZud3q8CkQJ9bh3CyAwYAiVC0"}).
		Expect(t).
		Status(200).
		End()
}
