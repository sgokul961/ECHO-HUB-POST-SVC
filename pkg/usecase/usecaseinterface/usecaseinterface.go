package interfacesU

type PostUseCaseInterface interface {
	FollowUser(following_id, follower_id int64) (bool, error)
}
