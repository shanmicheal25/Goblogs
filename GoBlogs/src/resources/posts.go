package models

import "time"

//Posts structure
type Posts struct {
	PostID      uint       `json:"post_id"`
	UserID      uint       `json:"user_id"`
	Text        string     `json:"text"`
	ImageURL    string     `json:"image_url"`
	CreateAt    time.Time  `json:"create_at,omitempty"`
	ShowPost    string     `json:"show_post"`
	CommentText string     `json:"comment_text"`
	Comments    []Comments `json:"comments"`
}
