package controllers

import (
	"strconv"

	u "GoBlogs/apiHelpers"

	ser "GoBlogs/services"

	"github.com/gin-gonic/gin"
)

//SearchBlogsReq struct
type SearchBlogsReq struct {
	UserID   uint `form:"user_id" json:"user_id"`
	Nextpage int  `form:"nextpage" json:"nextpage"`
}

// UsersReq ..
type BlogsReq struct {
	UserID uint   `json:"user_id"`
	Text   string `json:"text"`
}

// Blogsfilter ...
func Blogsfilter(c *gin.Context) {
	var blogsService ser.BlogsService
	var searchBlogsReq SearchBlogsReq
	c.ShouldBind(&searchBlogsReq)
	//page, _ := strconv.Atoi(c.DefaultQuery("page", searchProdReq.Nextpage))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", u.Limit))

	resp := blogsService.BlogsFilter(searchBlogsReq.Nextpage, limit)
	//return response using api helper
	u.Respond(c.Writer, resp)
}

// BlogPost ..
func BlogPost(c *gin.Context) {

	var blogService ser.BlogService

	var blogsReq BlogsReq
	c.BindJSON(&blogsReq)

	blogService.Posts.UserID = blogsReq.UserID
	blogService.Posts.Text = blogsReq.Text

	resp := blogService.BlogsPost()

	//return response using api helper
	u.Respond(c.Writer, resp)
}
