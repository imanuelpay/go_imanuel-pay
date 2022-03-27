package main

import (
	"mvc/config"
	"mvc/middlewares"
	"mvc/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
