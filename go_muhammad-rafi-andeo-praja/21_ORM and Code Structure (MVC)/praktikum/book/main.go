package main

import (
	"project-structure/config"
	"project-structure/routes"
)

func main() {
	config.InitDB()
	r := routes.New()
	r.Start(":8000")
}
