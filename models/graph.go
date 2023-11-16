package models

type Graph struct {
	UserID     uint64 `json:"user_id"`     // The user who made the follow request
	FolloweeID uint64 `json:"follower_id"` // The user who is being followed
	CreatedAt  string `json:"created_at"`
}
