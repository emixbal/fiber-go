package routers

import (
	"fiber-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("/ path") })
	app.Get("/pegawai", controllers.FetchAllPegawais)
	app.Put("/pegawai/:id", controllers.UpdatePegawai)

	return app
}
