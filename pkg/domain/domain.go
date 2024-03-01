package domain

type Follow struct {
	ID              int64 `json:"id" gorm:"primaryKey"`
	FollowingUserID int64 `json:"following_user_id" `
	FollowerUserID  int64 `json:"follower_user_id" `
}
