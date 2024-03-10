package domain

import (
	"time"
)

type Follow struct {
	ID              int64 `json:"id" gorm:"primaryKey"`
	FollowingUserID int64 `json:"following_user_id" `
	FollowerUserID  int64 `json:"follower_user_id" `
}

// Post represents the post table in the database
type Post struct {
	PostID        int64     `gorm:"primaryKey" json:"post_id"`
	UserID        int64     `json:"user_id"`
	Content       string    `json:"content" gorm:"not null"`
	ImageURL      string    `json:"image_url,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
	LikesCount    int64     `json:"likes_count" gorm:"default:0"`
	CommentsCount int64     `json:"comments_count"`
	Comments      []Comment `json:"comments,omitempty" gorm:"foreignKey:PostsID;constraint:OnDelete:CASCADE"`
}

// Like represents the like table in the database
type Like struct {
	LikeID    int64     `gorm:"primaryKey" json:"like_id"`
	PostsID   int64     `json:"posts_id"`
	UserID    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
	Post      Post      `json:"post" gorm:"foreignKey:PostsID;constraint:OnDelete:CASCADE"`
}

// Comment represents the comment table in the database
type Comment struct {
	CommentID int64     `gorm:"primaryKey" json:"comment_id"`
	PostsID   int64     `json:"posts_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content" gorm:"not null"`
	Timestamp time.Time `json:"timestamp"`
	Post      Post      `json:"post" gorm:"foreignKey:PostsID;constraint:OnDelete:CASCADE"`
}

// Tag represents the tag table in the database

type Tag struct {
	TagID   uint64 `gorm:"primaryKey" json:"tag_id"`
	TagName string `json:"tag_name" gorm:"not null;unique"`
}

// PostTag represents the post_tags table in the database
type PostTag struct {
	PostID uint64 `json:"post_id" gorm:"foreignKey:PostID"`
	TagID  uint64 `json:"tag_id" gorm:"foreignKey:TagID"`
}
