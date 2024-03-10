package interfacesR

import "github.com/sgokul961/echo-hub-post-svc/pkg/domain"

type PostRepoInterface interface {
	CreateFollowing(following_id, follower_id int64) (bool, error)
	FollowingExist(following_id int64) bool
	FollowerExist(follower_id int64) bool
	FollowRelationshipExists(following_id, follower_id int64) bool
	Unfollow(follower_id, following_id int64) error
	AddPost(post domain.Post) (int64, error)
	DeletePost(post_id int64, user_id int64) (int64, error)
	PostIdExist(post_id, user_id int64) bool
	LikePost(user_id, post_id int64) (int64, error)
	CheckForPostId(post_id int64) bool
	UpdatePost(post_id int64) bool
	UpdatePostDislike(post_id int64) bool

	AlredyLiked(postId, userId int64) bool
	DisLikePost(post_id, user_id int64) bool
	ChekIfLikeExist(post_id, user_id int64) bool
}
