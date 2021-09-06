package test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TestApplication struct {
	Router *gin.Engine
}

func TestEnvs(t *testing.T) {
	_, exists := os.LookupEnv("USER_TOKEN")
	if !exists {
		err := godotenv.Load(".env")
		if err != nil {
			t.Errorf(os.Environ())
		}
	}
}
