package router

import (
	"database/sql"
	"mawakif/internal/handlers"

	"github.com/gin-gonic/gin"
)

//Router is contains a handler type
type router struct {
	DB *sql.DB
	R  *gin.Engine
}

//New creates a new router instance
func New(db *sql.DB) router {
	return router{DB: db, R: gin.New()}
}

//Route routes the different requests
func (r router) Route() {
	//starts a new router
	r.R.GET("/api/subscribers", nil)
	r.R.GET("/api/checks", nil)
	r.R.GET("/api/space", nil)

	r.R.POST("api/plate", nil)
	r.R.POST("/api/subcribers/add", nil)

	r.R.PUT("/api/admin/ticket-duration", handlers.UpdateTicketDuration(r.DB))
}

//Run starts a listen and serve on a port
func (r router) Run(port string) error {
	return r.R.Run(port)
}
