package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors handles cors
func Cors(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		return
	}
	c.Next()
}
