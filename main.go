package main

import (
	"fiber-go/db"
	"fiber-go/routers"
)

func main() {
	app := routers.Init()
	db.Init()
	app.Listen(":3000")
}
