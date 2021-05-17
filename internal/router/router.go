package router

import (
	"mawakif/config"
	"mawakif/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Router is contains a handler type
type router struct {
	DB     *gorm.DB
	R      *gin.Engine
	CONFIG config.CONFIG
}

//New creates a new router instance
func New(cfg config.CONFIG, db *gorm.DB) router {
	return router{DB: db, //temperal
		R:      gin.New(),
		CONFIG: cfg} //config- temperal
}

//Route routes the different requests
func (r router) Route() {

	r.R.Use(Cors)
	//starts a new router
	r.R.GET("/api/subscribers", handlers.GetAllSubscribersHandler(r.DB))
	r.R.GET("/api/checks", handlers.GetAllChecksHandler(r.DB))
	r.R.GET("/api/packing-space", handlers.GetAllPackingSpaceHandler(r.DB))

	r.R.POST("/api/subcribers/add", handlers.AddUpdateSubscriberHandler(r.DB))
	r.R.POST("/api/checks/add", handlers.AddChecksHandler(r.DB))
	r.R.POST("/api/packing-space/add", handlers.AddPackingSpaceHandler(r.DB))

	r.R.PUT("/api/admin/ticket-duration", handlers.UpdateTicketDurationHandler(r.DB))
}

//Run starts a listen and serve on a port
func (r router) Run(port string) error {
	return r.R.Run(port)
}
