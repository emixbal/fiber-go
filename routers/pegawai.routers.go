package routers

import (
	"fiber-go/controllers"
	"fiber-go/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Pegawai(app *fiber.App) {
	pegawai := app.Group("/pegawai")

	// // define middleware
	// pegawai.Use(middlewares.ExampleMiddleware)
	// more routes
	pegawai.Get(
		"/",
		middlewares.IsAuthenticated, //middleware
		controllers.FetchAllPegawais,
	)
	pegawai.Put("/:id", controllers.UpdatePegawai)
}
