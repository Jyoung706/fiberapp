package routers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	UserRouter(app)
}
