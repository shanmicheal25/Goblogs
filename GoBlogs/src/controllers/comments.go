package controllers

import (
	"strconv"
	"time"

	u "GoBlogs/apiHelpers"

	ser "GoBlogs/services"

	"github.com/gin-gonic/gin"
)

//SearchBlogsReq struct
type SearchCommentsReq struct {
	CommentLevel uint `form:"comment_level" json:"comment_level"`
	CommentID    uint `form:"comment_id" json:"comment_id"`
	PostID       uint `form:"post_id" json:"post_id"`
	Nextpage     int  `form:"nextpage" json:"nextpage"`
}

// UsersReq ..
type CommentReq struct {
	UserID       uint   `json:"user_id"`
	CommentText  string `json:"comment_text"`
	ParentCommID uint   `json:"parent_comm_id"`
	PostID       uint   `json:"post_id"`
	CommentLevel uint   `json:"comment_level"`
	CommentID    uint   `json:"comment_id"`
}

// Commentsfilter ...
func Commentsfilter(c *gin.Context) {
	var commentsService ser.CommentsService
	var searchCommentsReq SearchCommentsReq
	c.ShouldBind(&searchCommentsReq)
	//page, _ := strconv.Atoi(c.DefaultQuery("page", searchProdReq.Nextpage))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", u.Limit))

	resp := commentsService.CommentsFilter(searchCommentsReq.PostID, searchCommentsReq.Nextpage, limit)
	//return response using api helper
	u.Respond(c.Writer, resp)
}

// SubCommentsfilter ...
func SubCommentsfilter(c *gin.Context) {
	var commentsService ser.CommentsService
	var searchCommentsReq SearchCommentsReq
	c.ShouldBind(&searchCommentsReq)
	//page, _ := strconv.Atoi(c.DefaultQuery("page", searchProdReq.Nextpage))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", u.Limit))

	resp := commentsService.SubCommentsFilter(searchCommentsReq.CommentID, searchCommentsReq.Nextpage, limit)
	//return response using api helper
	u.Respond(c.Writer, resp)
}

// BlogPost ..
func CommentPost(c *gin.Context) {

	var commService ser.CommentService

	var commReq CommentReq
	c.BindJSON(&commReq)

	commService.Comment.UserID = commReq.UserID
	commService.Comment.PostID = commReq.PostID
	commService.Comment.ParentCommID = commReq.ParentCommID
	commService.Comment.CommentText = commReq.CommentText
	commService.Comment.CommentLevel = commReq.CommentLevel + 1
	commService.Comment.CreateAt =  time.Now()
	commService.Comment.ShowComment = "Y"
	commService.Comment.ReplyCount = 0

	resp := commService.CommentPost()

	//return response using api helper
	u.Respond(c.Writer, resp)
}
