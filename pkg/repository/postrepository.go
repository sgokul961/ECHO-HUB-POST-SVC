package repository

import (
	"errors"
	"fmt"

	"github.com/sgokul961/echo-hub-post-svc/pkg/domain"
	interfacesR "github.com/sgokul961/echo-hub-post-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type PostDatabase struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) interfacesR.PostRepoInterface {
	return &PostDatabase{DB: db}
}

func (p *PostDatabase) FollowRelationshipExists(following_id, follower_id int64) bool {
	fmt.Println("following ,follower:", following_id, follower_id)

	query := `SELECT EXISTS (SELECT 1 FROM follows WHERE following_user_id = ? AND follower_user_id = ?)`

	var exists bool

	row := p.DB.Raw(query, following_id, follower_id).Scan(&exists)

	if row.Error != nil { // <- Corrected error handling
		return false
	}

	// Return true if the follow relationship exists, otherwise false
	return exists
}
func (p *PostDatabase) CreateFollowing(following_id, follower_id int64) (bool, error) {
	fmt.Println("following ,follower:", following_id, follower_id)

	query := `INSERT INTO follows (following_user_id,follower_user_id)VALUES (?,?)`
	err := p.DB.Exec(query, following_id, follower_id).Error

	if err != nil {
		return false, errors.New("no rows were affected")

	}
	return true, nil
}

// not using this 2 repo functions
func (p *PostDatabase) FollowingExist(following_id int64) bool {
	var count int64

	query := `SELECT COUNT(*) FROM follows WHERE following_user_id = ?`

	err := p.DB.Raw(query, following_id).Scan(&count).Error

	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false

}

func (p *PostDatabase) FollowerExist(follower_id int64) bool {
	var count int64

	query := `SELECT COUNT(*) FROM follows WHERE follower_user_id=?`

	err := p.DB.Raw(query, follower_id).Scan(&count)

	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

// unfollow user
func (p *PostDatabase) Unfollow(following_id, follower_id int64) error {
	query := `DELETE FROM follows WHERE following_user_id= ? AND follower_user_id=?`

	result := p.DB.Exec(query, following_id, follower_id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// upload post
func (p *PostDatabase) AddPost(post domain.Post) (int64, error) {

	var userID int64

	query := `INSERT INTO posts (user_id, content, image_url,timestamp) VALUES (?, ?, ?, ?) RETURNING user_id`

	res := p.DB.Raw(query, post.UserID, post.Content, post.ImageURL, post.Timestamp).Scan(&userID)
	if res.Error != nil {
		return 0, res.Error
	}

	return userID, nil

}
func (u *PostDatabase) DeletePost(post_id int64, user_id int64) (int64, error) {

	query := `DELETE FROM posts WHERE post_id =? AND user_id = ?`

	err := u.DB.Exec(query, post_id, user_id).Error
	if err != nil {
		return 0, errors.New("database error")
	}

	return post_id, nil
}
func (p *PostDatabase) PostIdExist(post_id, user_id int64) bool {
	query := `SELECT EXISTS (SELECT 1 FROM posts WHERE post_id = ? AND user_id = ?)`
	var exists bool

	err := p.DB.Raw(query, post_id, user_id).Scan(&exists).Error
	if err != nil {
		return false
	}
	return exists
}
func (p *PostDatabase) CheckForPostId(post_id int64) bool {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM posts WHERE post_id = ?)`

	if err := p.DB.Raw(query, post_id).Scan(&exists).Error; err != nil {
		return false
	}
	return exists
}
func (p *PostDatabase) LikePost(user_id, post_id int64) (int64, error) {

	likeQuery := `INSERT INTO likes (posts_id,user_id,timestamp) VALUES(?,?,NOW())`

	if err := p.DB.Exec(likeQuery, post_id, user_id).Error; err != nil {
		return 0, err
	}

	return post_id, nil

}
func (p *PostDatabase) UpdatePost(post_id int64) bool {

	// Update the likes_count in the posts table
	updateQuery := `UPDATE posts SET likes_count = likes_count + 1 WHERE  post_id = ?`
	if err := p.DB.Exec(updateQuery, post_id).Error; err != nil {
		return false
	}
	return true

}

// impliment this
func (p *PostDatabase) UpdatePostDislike(post_id int64) bool {
	query := `UPDATE posts SET likes_count=likes_count-1 WHERE post_id =?`
	if err := p.DB.Exec(query, post_id).Error; err != nil {
		return false
	}
	return true

}

func (p *PostDatabase) AlredyLiked(postId, userId int64) bool {

	query := `SELECT EXISTS(SELECT 1 from likes WHERE posts_id = ? AND user_id = ?)`
	var exists bool

	err := p.DB.Raw(query, postId, userId).Scan(&exists).Error
	if err != nil {
		return false
	}
	return exists

}

// dislike post
func (p *PostDatabase) DisLikePost(post_id, user_id int64) bool {

	query := `DELETE FROM likes WHERE posts_id = ? AND user_id = ?`

	err := p.DB.Exec(query, post_id, user_id).Error

	if err != nil {
		fmt.Println("error executing query:", err)
		return false
	}

	return true
}

func (p *PostDatabase) ChekIfLikeExist(post_id, user_id int64) bool {
	fmt.Println("post and user ", post_id, user_id)

	query := `SELECT EXISTS (SELECT 1 FROM likes WHERE posts_id= ? AND user_id = ?)`

	var exist bool
	fmt.Println("post and user id ", post_id, user_id)

	err := p.DB.Raw(query, post_id, user_id).Scan(&exist).Error
	if err != nil {
		fmt.Println("error executing query:", err)

		return false
	}
	fmt.Println("exist ", exist)
	return exist

}
