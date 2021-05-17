package main

import (
	"log"
	"mawakif/config"
	"mawakif/internal/handlers"
	"mawakif/internal/router"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, cancelFunc, er := handlers.Connect(handlers.ConnectionString(cfg))
	if er != nil {
		log.Printf("failed to connect to database :%v\n", er.Error())
		return
	}
	defer cancelFunc()

	myRouter := router.New(cfg, db)
	myRouter.Route()
	myRouter.Run(cfg.PORT)
}
