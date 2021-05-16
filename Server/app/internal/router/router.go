package router

import (
	"database/sql"
	"mawakif/config"
	"mawakif/internal/handlers"

	"github.com/gin-gonic/gin"
)

//Router is contains a handler type
type router struct {
	DB     *sql.DB
	R      *gin.Engine
	CONFIG config.CONFIG
}

//New creates a new router instance
func New(cfg config.CONFIG) router {
	return router{DB: nil, //temperal
		R:      gin.New(),
		CONFIG: cfg} //config- temperal
}

//Route routes the different requests
func (r router) Route() {
	//starts a new router
	r.R.GET("/api/subscribers", nil)
	r.R.GET("/api/checks", nil)
	r.R.GET("/api/space", nil)

	r.R.POST("api/plate", nil)
	r.R.POST("/api/subcribers/add", nil)

	r.R.PUT("/api/admin/ticket-duration", handlers.UpdateTicketDuration(r.CONFIG))
}

//Run starts a listen and serve on a port
func (r router) Run(port string) error {
	return r.R.Run(port)
}
