package handlers

import (
	"encoding/json"
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

//AddPackingSpaceHandler adds new packing spaces
func AddPackingSpaceHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		prksp := database.PackingSpace{}

		err := json.NewDecoder(c.Request.Body).Decode(&prksp)
		if err != nil {
			httperror.Default(err).ReplyBadRequest(c.Writer)
			return
		}


		er := prksp.AddUpdate(db)
		if er != nil {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(prksp).ReplyOK(c.Writer)
	}
}
