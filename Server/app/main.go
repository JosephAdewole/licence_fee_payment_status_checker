package main

import (
	"log"
	"mawakif/config"
	"mawakif/internal/router"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Println(err.Error())
		return
	}
	myRouter := router.New(cfg)
	myRouter.Route()
	myRouter.Run(":3000")
}
