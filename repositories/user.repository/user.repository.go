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

func UpdateUser(id int, user userModel.User) string {
	for i, u := range userList {
		if u.ID == id {
			userList[i] = user
		}
	}
	return "User has been updated"
}

func DeleteUser(id int) string {
	for i, user := range userList {
		if user.ID == id {
			userList = append(userList[:i], userList[i+1:]...)
		}
	}
	return "User has been deleted"
}
