package interfacesU

import "github.com/sgokul961/echo-hub-post-svc/pkg/models"

type PostUseCaseInterface interface {
	FollowUser(following_id, follower_id int64) (bool, error)
	UnfollowUser(following_id, follower_id int64) (int64, error)
	AddPost(upload models.AddPost) (int64, error)
	DeletePost(post_id, user_id int64) (int64, error)
	LikePost(post_id, user_id int64) (int64, error)
	DisLikepost(user_id, post_id int64) (bool, error)
	AddComment(comment models.AddComent) (int64, error)
	GetComment(post_id int64) ([]string, error)
	DeleteComment(postID, commentID, UserID int64) (int64, bool)
	GetPostOwner(postid int64) (int64, error)
}
