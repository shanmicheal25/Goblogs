package services

import (
	u "GoBlogs/apiHelpers"
	"GoBlogs/models"
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
)

//CommentsService struct
type CommentsService struct {
	Comments []models.Comments
}

//CommentService struct
type CommentService struct {
	Comment models.Comments
}

//CommentsFilter function returns the list
func (comService CommentsService) CommentsFilter(post_id uint, page int, limit int) map[string]interface{} {

	commArr := comService.Comments
	paginator := pagination.Paging(&pagination.Param{
		DB:      models.GetDB().Where("post_id =  ? and comment_level = 1 and show_comment = 'Y'", post_id),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"post_id desc"},
		ShowSQL: true,
	}, &commArr)

	if paginator == nil {
		response := u.Message(1, "Invalid request")
		return response
	}

	response := u.Message(0, "Post Comments Result ")
	response["commentsData"] = paginator
	return response
}

//SubCommentsFilter function returns the list
func (comService CommentsService) SubCommentsFilter(comment_id uint, page int, limit int) map[string]interface{} {

	commArr := comService.Comments
	paginator := pagination.Paging(&pagination.Param{
		DB:      models.GetDB().Where("parent_comm_id =  ?  and show_comment = 'Y'", comment_id),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"post_id desc"},
		ShowSQL: true,
	}, &commArr)

	if paginator == nil {
		response := u.Message(1, "Invalid request")
		return response
	}

	response := u.Message(0, "Post Comments Result ")
	response["commentsData"] = paginator
	return response
}

//CommentsPost record .
func (comm CommentService) CommentPost() map[string]interface{} {

	comment := comm.Comment
	response := u.Message(200, "Comment Posted successfully")

	err := models.GetDB().Create(&comment).Error
	if err != nil {
		fmt.Println("Error ", err)
		response = u.Message(1, "Invalid request")
	}

	var commModel models.Comments
	models.GetDB().Where("comment_id = ?", comment.ParentCommID).First(&commModel)

	count := commModel.ReplyCount + 1
	models.GetDB().Table("comments").Where("comment_id = ?", comment.ParentCommID).Update("reply_count", count)

	response["commentsData"] = comment
	return response

}
