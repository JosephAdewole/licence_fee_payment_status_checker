package handlers

import (
	"encoding/json"
	"errors"
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetAllPackingSpaceHandler gets alls the  checks done by the robot
func GetAllPackingSpaceHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		prksp := database.PackingSpace{}
		prksps, er := prksp.GetAll(db)
		if er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(prksps).ReplyOK(c.Writer)
	}
}

type addPackingSpaceRequest struct {
	ID          uint   `json:"id"`
	Designation string `json:"designation"`
}

func (a addPackingSpaceRequest) validate() error {
	if a.ID == 0 {
		return errors.New("id is required")
	}

	if a.Designation == "" {
		return errors.New("designation is required")
	}

	return nil
}

//AddPackingSpaceHandler adds new packing spaces
func AddPackingSpaceHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req addPackingSpaceRequest

		err := json.NewDecoder(c.Request.Body).Decode(&req)
		if err != nil {
			httperror.Default(err).ReplyBadRequest(c.Writer)
			return
		}

		if er := req.validate(); er != nil {
			httperror.Default(err).ReplyBadRequest(c.Writer)
			return
		}

		pks := database.PackingSpace{
			ID:          req.ID,
			Designation: req.Designation,
		}

		er := pks.AddUpdate(db)
		if er != nil {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(pks).ReplyOK(c.Writer)
	}
}
