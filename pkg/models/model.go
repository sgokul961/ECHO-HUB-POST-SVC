package models

import "time"

type AddPost struct {
	PostID    int64     `gorm:"primaryKey" json:"post_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content" gorm:"not null"`
	ImageURL  string    `json:"image_url,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
type AddComent struct {
	PostsID   int64     `json:"posts_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content" gorm:"not null"`
	Timestamp time.Time `json:"timestamp"`
}
type Notification struct {
	Topic   string `json:"topic"`
	UserID  int64  `json:"user_id"`
	PostsID int64  `json:"post_id"`
	Message string `json:"message"`
}
