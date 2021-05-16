package main

import (
	"log"
	"mawakif/config"
	"mawakif/internal/router"
	"os"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Println(err.Error())
		return
	}
	myRouter := router.New(cfg)
	myRouter.Route()
	port := ":" + os.Getenv("PORT")
	myRouter.Run(port)
}
