package main

import (
	"log"
	"mawakif/config"
	"mawakif/internal/database"
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

func init() {
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

	db.AutoMigrate(&database.Check{})
	db.AutoMigrate(&database.Config{})
	db.AutoMigrate(&database.PackingSpace{})
	db.AutoMigrate(&database.Subscriber{})
}
