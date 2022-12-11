package main

import (
	"kitabisa_api_golang/db"
	"kitabisa_api_golang/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":6969"))
}
