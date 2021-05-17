package main

import (
	"log"
	"mawakif/config"
	"mawakif/internal/handlers"
	"mawakif/internal/router"
	"os"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, cancelFunc, er := handlers.Connect(handlers.ConnectionString(cfg))
	if er != nil {
		log.Println(er.Error())
		return
	}
	defer cancelFunc()

	myRouter := router.New(cfg, db)
	myRouter.Route()
	port := ":" + os.Getenv("PORT")
	myRouter.Run(port)
}
