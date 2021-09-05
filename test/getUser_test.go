package test

import (
	c "goychatapp/controllers"
	jwt "goychatapp/lib"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

var token string

func TestGetUser(t *testing.T) {
	if os.Getenv("TEST_ENV") == "true" {
		token = os.Getenv("USER_TEST_TOKEN")
	} else {
		token = "test"
	}
	r := gin.Default()
	r.GET("/v1/user", jwt.UserAdminAuth, c.GetProfile)
	testRoute := &TestApplication{Router: r}
	apitest.New().
		Handler(testRoute.Router).
		Get("/v1/user").
		Headers(map[string]string{"Authorization": token}).
		Expect(t).
		Status(200).
		End()
}
