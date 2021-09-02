package main

import (
	"github.com/gofiber/fiber/v2"

	"fiber-go/db"
	"fiber-go/routers"
)

func main() {
	app := fiber.New()

	routers.Init(app)
	db.Init()

	app.Listen(":3000")
}
