package userService

import (
	"fiberapp/models/userModel"
	userRepository "fiberapp/repositories/user.repository"
)

func CreateUser(userForm userModel.User) string {
	return userRepository.CreateUser(userForm)
}

func GetUser(id int) userModel.User {
	return userRepository.GetUser(id)
}
