package routers

import (
	"fiber-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App) {
	login := app.Group("/auth")
	login.Post("/register", controllers.Register)
	login.Post("/login", controllers.Login)
}
