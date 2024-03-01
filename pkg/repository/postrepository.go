package repository

import (
	"errors"
	"fmt"

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
