package handlers

import (
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"

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
