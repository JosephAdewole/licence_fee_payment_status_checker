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
	PlateNumber  string    `json:"plate_number"`
	PlateSpaceID string    `json:"packing_space_id"`
	CurrentTime  time.Time `json:"current_time"`
	IsEmpty      bool      `json:"is_empty"`
}

//AddChecksHandler accepts check log and stores to database
func AddChecksHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req addChecksRequest
		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		er := db.Model(&database.Check{}).Create(&req).Error
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		if req.IsEmpty {
			return
		}

		er = db.Model(&database.Subscriber{}).Where("plate_number=?", req.PlateNumber).FirstOrCreate(&database.Subscriber{
			PlateNumber: req.PlateNumber}).Error
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}
	}
}
