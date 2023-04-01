package main

import (
	"project/config"
	"project/routes"
)

func main() {
	config.InitDB()
	r := routes.New()
	r.Logger.Fatal(r.Start(":8000"))
}
