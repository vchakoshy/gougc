package follow

type FollowRequest struct {
	LeaderID   uint `json:"leader_id" binding:"required"`
	FollowerID uint `json:"follower_id" binding:"required"`
}
