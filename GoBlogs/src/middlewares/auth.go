package middlewares

import (
	//"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	hel "GoBlogs/apiHelpers"
)

var identityKey = "id"

type login struct {
	UserID  string `form:"userId" json:"userId" binding:"required"`
	Appname string `form:"appname" json:"appname" binding:"required"`
}

// User demo
type User struct {
	UserID string `form:"userId" json:"userId" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

/*
UserMiddlewares function to add auth
*/
func UserMiddlewares() gin.HandlerFunc {

	return func(c *gin.Context) {

		err := hel.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}
