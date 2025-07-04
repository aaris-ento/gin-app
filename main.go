package main

import (
	"gin-app/bootstrap"
	"gin-app/routes"
)

func main() {
	app := bootstrap.Init()
	r := routes.SetupRoutes(app)

	r.Run(":8080")
}
