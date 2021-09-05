package test

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steinfletcher/apitest"
)

func TestGetUser(t *testing.T) {
	_, exists := os.LookupEnv("TEST_ENV")
	if !exists {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Baby")
			log.Fatal("Error load the .env file")
		}
	}
	r := gin.Default()
	r.GET("/v1/user", jwt.UserAdminAuth, c.GetProfile)
	testRoute := &TestApplication{Router: r}
	apitest.New().
		Handler(testRoute.Router).
		Get("/v1/user").
		Headers(map[string]string{"Authorization": os.Getenv("USER_TEST_TOKEN")}).
		Expect(t).
		Status(200).
		End()
}
