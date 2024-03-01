package routers

import (
	userController "fiberapp/controllers/user.controller"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	app.Post("/user", userController.CreateUser)
	app.Get("/user/:id", userController.GetUser)
}
