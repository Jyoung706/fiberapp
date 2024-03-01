package userController

import (
	"fiberapp/models/userModel"
	userService "fiberapp/services/user.service"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := userModel.User{}
	c.BodyParser(&user)
	response := userService.CreateUser(user)

	return c.SendString(response)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := userService.GetUser(id)

	return c.JSON(user)
}
