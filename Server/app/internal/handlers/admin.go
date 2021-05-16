package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mawakif/config"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type updateTicketDurationRequest struct {
	Duration int `json:"ticket_duration" validate:"required"`
}

func (u updateTicketDurationRequest) validateUpdateTicketDurationRequest() error {
	if u.Duration == 0 {
		return errors.New("duration is required")
	}

	return nil
}

func (u updateTicketDurationRequest) string() string {
	return strconv.Itoa(u.Duration)
}

//UpdateTicketDuration updates the duration of tickets
func UpdateTicketDuration(cfg config.CONFIG) func(c *gin.Context) {

	configString, err := readFile(cfg)
	if err != nil {
		return func(c *gin.Context) {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
		}
	}

	return updateTicketDurationHandler(configString, cfg)

}

//returns a handler func for updating ticket duration
func updateTicketDurationHandler(configString string, cfg config.CONFIG) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req updateTicketDurationRequest

		if er := json.NewDecoder(c.Request.Body).Decode(&req); er != nil {
			httperror.Default(errors.Wrap(er, "unable to read request body")).ReplyBadRequest(c.Writer)
			return
		}

		if er := req.validateUpdateTicketDurationRequest(); er != nil {
			httperror.Default(er).ReplyBadRequest(c.Writer)
			return
		}

		newStr := strings.Replace(configString, cfg.TicketDuration, req.string(), -1)
		if er := writeFile(cfg, newStr); er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		msg := fmt.Sprintf("ticket duration updated to %v hrs", req.Duration)
		httpresp.New(true, 200, msg, nil, nil).ReplyOK(c.Writer)
	}
}

func readFile(cfg config.CONFIG) (string, error) {

	fs, err := os.OpenFile(cfg.ConfigFileName, os.O_RDWR, 777)
	if err != nil {
		return "", err
	}
	defer fs.Close()

	data, err := ioutil.ReadAll(fs)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func writeFile(cfg config.CONFIG, content string) error {
	fs, err := os.OpenFile(cfg.ConfigFileName, os.O_RDWR, 777)
	if err != nil {
		return err
	}
	defer fs.Close()

	_, er := fs.WriteString(content)
	if er != nil {
		return er
	}

	return nil
}
