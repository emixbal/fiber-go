package routers

import (
	"fiber-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/hai", func(c *fiber.Ctx) error {
		return c.SendString("hai")
	})

	app.Get("/pegawai", controllers.FetchAllPegawai)

	return app
}
