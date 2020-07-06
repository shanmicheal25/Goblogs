package models

import "time"

//Comments structure
type Comments struct {
	CommentID    uint      `json:"comment_id"`
	UserID       uint      `json:"user_id"`
	CommentText  string    `json:"comment_text"`
	ParentCommID uint      `json:"parent_comm_id"`
	CommentLevel uint      `json:"comment_level"`
	ReplyCount   uint      `json:"reply_count"`
	CreateAt     time.Time `json:"create_at,omitempty"`
	ShowComment  string    `json:"show_comment"`
}
