package main

import (
	"mawakif/internal/router"
)

func main() {

	myRouter := router.New(nil)
	myRouter.Route()
	myRouter.Run(":3000")
}
