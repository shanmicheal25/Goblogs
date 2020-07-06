package models

import (
	"time"
)

//Posts structure
type Posts struct {
	PostID   uint       `gorm:"primary_key;" json:"post_id"`
	UserID   uint       `gorm:"type:int"  json:"user_id"`
	Text     string     `gorm:"type:varchar(5000)" json:"text"`
	ImageURL string     `gorm:"type:varchar(50)" json:"image_url"`
	CreateAt time.Time  `gorm:"type:datetime" json:"create_at"`
	ShowPost string     `gorm:"type:varchar(1)" json:"show_post"`
	Comments []Comments `gorm:"foreignkey:PostID" json:"comments"`
}

//TableName return name of database table
func (p *Posts) TableName() string {
	return "posts"
}

// func (post Posts) ToString() string {
// 	return fmt.Sprintf("id: %d\nname: %s", post.PostID, post.Text)
// }
