package models

import "time"

type AddPost struct {
	PostID    int64     `gorm:"primaryKey" json:"post_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content" gorm:"not null"`
	ImageURL  string    `json:"image_url,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
