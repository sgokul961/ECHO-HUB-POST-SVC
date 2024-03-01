package interfacesR

type PostRepoInterface interface {
	CreateFollowing(following_id, follower_id int64) (bool, error)
	FollowingExist(following_id int64) bool
	FollowerExist(follower_id int64) bool
	FollowRelationshipExists(following_id, follower_id int64) bool
	Unfollow(follower_id, following_id int64) error
}
