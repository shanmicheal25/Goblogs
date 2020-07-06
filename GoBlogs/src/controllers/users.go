package controllers

import (
	u "GoBlogs/apiHelpers"

	ser "GoBlogs/services"

	"github.com/gin-gonic/gin"
)

// UsersReq ..
type UserReq struct {
	UserID   uint   `form:"user_id" json:"user_id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Mobile   string `form:"mobile" json:"mobile"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

// UserRegister ..
func UserRegister(c *gin.Context) {

	var userService ser.UserService

	var userReq UserReq
	c.BindJSON(&userReq)

	userService.Users.UserID = userReq.UserID
	userService.Users.Name = userReq.Name
	userService.Users.Email = userReq.Email
	userService.Users.Mobile = userReq.Mobile
	userService.Users.Password = userReq.Password
	userService.Users.Address = userReq.Address

	resp := userService.UserRegister()

	//return response using api helper
	u.Respond(c.Writer, resp)
}
