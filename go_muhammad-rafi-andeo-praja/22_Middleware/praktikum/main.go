package main

import (
	"project-structure/middlewares"
	"project-structure/routes"
)

func main() {
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
