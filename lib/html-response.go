package lib

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTMLResponse(html string, statusCode int, c *gin.Context) *gin.Context {
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(html))
	return c
}
func HTMLErrorResponse(e error, statusCode int, html string, c *gin.Context) *gin.Context {
	if e == sql.ErrNoRows {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(html))
		return c
	}
	fmt.Println(e.Error())
	errorString := fmt.Sprintf("<b>%s</b>", e.Error())
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(errorString))
	return c
}
