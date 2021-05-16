package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type updateTicketDurationRequest struct {
	Duration time.Time
}

//UpdateTicketDuration updates the duration of tickets
func UpdateTicketDuration(db *sql.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var req updateTicketDurationRequest
		if er := c.BindJSON(&req); er != nil {
			c.Writer.WriteHeader(http.StatusUnprocessableEntity)
			c.Writer.Write([]byte(er.Error()))
			return
		}
		c.Bind(map[string]interface{}{
			"api": "/api/admin/ticket-duration",
		})
	}

}
