package main

import (
	"app/config"
	"app/routes"
)

func main() {
	e := routes.New()
	config.ConnectDB()
	e.Logger.Fatal(e.Start(":8000"))
}

