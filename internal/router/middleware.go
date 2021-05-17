package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors handles cors
func Cors(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")

	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

	if c.Request.Method == http.MethodOptions {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}
	c.Next()
}
