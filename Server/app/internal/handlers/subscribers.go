package handlers

import (
	"database/sql"
	"encoding/json"
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"time"

	"github.com/gin-gonic/gin"
)

type subscribersRequest struct {
	PlateNumber string    `json:"plate_number"`
	Status      bool      `json:"status"`
	CurrentTime time.Time `json:"current_time"`
}

//AddUpdateSubscriberHandler add a new subscriber or updates an existing one
func AddUpdateSubscriberHandler(db *sql.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var req subscribersRequest
		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		var strtTime time.Time
		if req.Status == true {
			strtTime = req.CurrentTime
		} else {
			strtTime = time.Time{}
		}

		subscriber := database.Subscriber{
			PlateNumber: req.PlateNumber,
			Status:      req.Status,
			StartTime:   strtTime}

		if er := subscriber.Add(db); er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(subscriber).ReplyCreated(c.Writer)
	}

}

//GetAllSubscribersHandler gets all subscribers from database
func GetAllSubscribersHandler(db *sql.DB) func(c *gin.Context) {

	return func(c *gin.Context) {

		subscriber := database.Subscriber{}
		subscribers, er := subscriber.GetAll(db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(subscribers).ReplyOK(c.Writer)
	}

}
