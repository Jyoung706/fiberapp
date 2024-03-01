package userRepository

import (
	"fiberapp/models/userModel"
)

var userList []userModel.User

func CreateUser(user userModel.User) string {
	userList = append(userList, user)
	return "User has been created"
}

func GetUser(id int) userModel.User {
	for _, user := range userList {
		if user.ID == id {
			return user
		}
	}
	return userModel.User{}
}
