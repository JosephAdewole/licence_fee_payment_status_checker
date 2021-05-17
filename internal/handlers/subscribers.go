package handlers

import (
	"encoding/json"
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type subscribersRequest struct {
	PlateNumber string    `json:"plate_number" validate:"require"`
	Status      bool      `json:"status" validate:"required"`
	CurrentTime time.Time `json:"current_time" validate:"required"`
}

//AddUpdateSubscriberHandler add a new subscriber or updates an existing one
func AddUpdateSubscriberHandler(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var req subscribersRequest
		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		if er := validator.New().Struct(req); er != nil {
			httperror.Default(er.(validator.ValidationErrors)).ReplyBadRequest(c.Writer)
			return
		}

		subscriber, er := addUpdateSubscribers(req, db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(subscriber).ReplyCreated(c.Writer)
	}

}

func addUpdateSubscribers(req subscribersRequest, db *gorm.DB) (database.Subscriber, error) {

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

	if er := subscriber.AddUpdate(db); er != nil {
		return database.Subscriber{}, er
	}

	return subscriber, nil
}

//GetAllSubscribersHandler gets all subscribers from database
func GetAllSubscribersHandler(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {

		subscriber := database.Subscriber{}
		subscribers, er := subscriber.GetAll(db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		ticketDuration := database.Config{Name: "TICKET_DURATION"}
		_ = ticketDuration.Get(db)

		tme, _ := strconv.Atoi(ticketDuration.Value)

		for i, v := range subscribers {
			k := time.Duration(tme) * time.Hour
			if time.Now().After(v.StartTime.Add(k)) {
				subscribers[i].Status = false
				subscriber.AddUpdate(db)
			}

		}

		httpresp.Default(subscribers).ReplyOK(c.Writer)
	}

}
