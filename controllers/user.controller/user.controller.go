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

	if user.ID == 0 {
		return c.Status(404).SendString("No user found with ID")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := userModel.User{}
	c.BodyParser(&user)
	response := userService.UpdateUser(id, user)
	return c.SendString(response)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	response := userService.DeleteUser(id)
	return c.SendString(response)
}
