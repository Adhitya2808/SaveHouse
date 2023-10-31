package main

import (
	"SaveHouse/config"
	"SaveHouse/routes"
)

func main() {
	e := routes.New()
	config.ConnectDB()
	e.Logger.Fatal(e.Start(":8000"))
}

