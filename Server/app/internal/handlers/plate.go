package handlers

import (
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type plateNumberRequest struct {
	PlateNumber    string    `json:"plate_number"`
	PackingSpaceID int       `json:"packing_space_id"`
	CurrentTime    time.Time `json:"current_time"`
	IsEmpty        bool      `json:"is_empty"`
}

//AddPlateNumberHandler logs robot activities
//and also adds subscriber to database
func AddPlateNumberHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req plateNumberRequest
		if er := c.BindJSON(&req); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		chk := database.Check{PlateNumber: req.PlateNumber,
			PackingSpaceID: req.PackingSpaceID,
			CreatedAt:      req.CurrentTime,
			IsEmpty:        req.IsEmpty}

		if er := chk.Add(db); er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}
		stat := true
		curTime := time.Now()

		input := subscribersRequest{PlateNumber: req.PlateNumber,
			Status:      stat,
			CurrentTime: curTime}

		subscriber, er := addUpdateSubscribers(input, db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(subscriber).ReplyCreated(c.Writer)

	}
}
