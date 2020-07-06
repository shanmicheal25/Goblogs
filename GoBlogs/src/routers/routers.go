package routers

import (
	"fmt"

	"GoBlogs/middlewares"

	apiController "GoBlogs/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to storage folder
	r.Static("/storage", "storage")

	//Giving access to template folder
	r.Static("/templates", "templates")
	r.LoadHTMLGlob("templates/*")

	r.Use(CORS())

	fmt.Printf("Test Routers part ")

	//API route for version 1
	v1 := r.Group("/api/v1")

	v1.POST("user", apiController.UserRegister)

	v1.GET("blogs", apiController.Blogsfilter)
	v1.POST("blogs", apiController.BlogPost)

	v1.GET("comments", apiController.Commentsfilter)
	v1.POST("comments", apiController.CommentPost)
	v1.GET("subcomments", apiController.SubCommentsfilter)

	//If you want to pass your route through specific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{

	}

	return r

}

// CORS ..
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
