package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mawakif/internal/database"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type updateTicketDurationRequest struct {
	TicketDuration int `json:"ticket_duration" validate:"required"`
}

func (u updateTicketDurationRequest) validateUpdateTicketDurationRequest() error {
	if u.TicketDuration == 0 {
		return errors.New("duration is required")
	}

	return nil
}

func (u updateTicketDurationRequest) string() string {
	return strconv.Itoa(u.TicketDuration)
}

//UpdateTicketDurationHandler returns a handler func for updating ticket duration
func UpdateTicketDurationHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req updateTicketDurationRequest

		//read request body
		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(errors.Wrap(er, "unable to read request body")).ReplyBadRequest(c.Writer)
			return
		}

		//validate request
		if er := req.validateUpdateTicketDurationRequest(); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		//update database
		conf := database.Config{Name: "TICKET_DURATION", Value: req.string()}
		if err := conf.Update(db); err == sql.ErrNoRows {
			if er := conf.Add(db); er != nil {
				httperror.Default(err).ReplyInternalServerError(c.Writer)
				return
			}
		} else if err != nil {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
			return
		}

		msg := fmt.Sprintf("%v updated to %v hrs", conf.Name, conf.Value)
		httpresp.New(true, 200, msg, nil, nil).ReplyOK(c.Writer)
	}
}
