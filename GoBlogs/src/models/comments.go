package models

import "time"

//Comments structure
type Comments struct {
	CommentID    uint      `gorm:"primary_key;" json:"comment_id"`
	PostID       uint      `gorm:"type:int" json:"post_id"`
	UserID       uint      `gorm:"type:int"  json:"user_id"`
	CommentText  string    `gorm:"type:varchar(5000)" json:"comment_text"`
	ParentCommID uint      `gorm:"type:int" json:"parent_comm_id"`
	CommentLevel uint      `gorm:"type:int" json:"comment_level"`
	ReplyCount   uint      `gorm:"type:int" json:"reply_count"`
	CreateAt     time.Time `gorm:"type:datetime" json:"create_at"`
	ShowComment  string    `gorm:"type:varchar(1)" json:"show_comment"`
	//Posts        Posts
}

//TableName return name of database table
func (c *Comments) TableName() string {
	return "comments"
}
