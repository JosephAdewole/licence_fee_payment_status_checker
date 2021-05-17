package handlers

import (
	"encoding/json"
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetAllChecksHandler gets alls the  checks done by the robot
func GetAllChecksHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		chk := database.Check{}
		chks, er := chk.GetAll(db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(chks).ReplyOK(c.Writer)
	}
}

type addChecksRequest struct {
	PlateNumber    string    `json:"plate_number"`
	PackingSpaceID uint      `json:"packing_space_id"`
	CurrentTime    time.Time `json:"current_time"`
	IsEmpty        bool      `json:"is_empty"`
}

//AddChecksHandler accepts check log and stores to database
func AddChecksHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req addChecksRequest
		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		ck := database.Check{
			PlateNumber:    req.PlateNumber,
			PackingSpaceID: req.PackingSpaceID,
			CreatedAt:      req.CurrentTime,
			IsEmpty:        req.IsEmpty,
		}
		er := ck.AddUpdate(db)
		if er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		if req.IsEmpty {
			return
		}

		sub := database.Subscriber{
			PlateNumber: req.PlateNumber,
		}

		err := sub.AddUpdate(db)

		if err != nil {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default([]interface{}{ck, sub}).ReplyCreated(c.Writer)

	}
}
