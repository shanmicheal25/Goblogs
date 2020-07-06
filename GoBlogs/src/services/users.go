package services

import (
	u "GoBlogs/apiHelpers"
	"GoBlogs/models"
	"fmt"
)

//UserService struct
type UserService struct {
	Users models.Users
}

// UserRegister record .
func (user UserService) UserRegister() map[string]interface{} {

	userData := user.Users

	err := models.GetDB().Create(&userData).Error
	if err != nil {
		fmt.Println("Error ", err)
		response := u.Message(1, "Invalid request")
		return response
	}

	response := u.Message(200, "User Created successfully")
	response["UserData"] = userData
	return response

}
