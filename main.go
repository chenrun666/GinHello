package main

import (
	_ "GinHello/docs"
	"GinHello/initRouter"
	)

// @host localhost:8080
func main() {
	router := initRouter.SetupRouter()

	router.Run()
}
