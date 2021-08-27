package main

import "fiber-go/routers"

func main() {
	app := routers.Init()
	app.Listen(":3000")
}
