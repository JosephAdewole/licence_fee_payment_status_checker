package handlers

import (
	"database/sql"
	"io/ioutil"
	"mawakif/config"
	"mawakif/pkg/httperror"
	"mawakif/pkg/httpresp"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type updateTicketDurationRequest struct {
	Duration string `json:"ticket_duration"`
}

//UpdateTicketDuration updates the duration of tickets
func UpdateTicketDuration(db *sql.DB, cfg config.CONFIG) func(c *gin.Context) {

	configString, err := readfile(cfg)
	if err != nil {
		return func(c *gin.Context) {
			httperror.Default(err).ReplyInternalServerError(c.Writer)
		}
	}

	return updateTicketDurationHandler(configString, cfg)

}

func updateTicketDurationHandler(configString string, cfg config.CONFIG) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req updateTicketDurationRequest
		if er := c.BindJSON(&req); er != nil {
			c.Writer.WriteHeader(http.StatusUnprocessableEntity)
			c.Writer.Write([]byte(er.Error()))
			return
		}

		newStr := strings.Replace(configString, cfg.TicketDuration, req.Duration, -1)
		if er := writeFile(cfg, newStr); er != nil {
			httperror.Default(er).ReplyInternalServerError(c.Writer)
			return
		}

		httpresp.Default(nil).ReplyOK(c.Writer)
	}
}

func readfile(cfg config.CONFIG) (string, error) {

	fs, err := os.OpenFile("app.config", 2, 777)
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
	fs, err := os.OpenFile("app.config", 2, 777)
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
