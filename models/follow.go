package models

import "time"

type Follow struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	LeaderID   uint      `json:"leader_id" gorm:"index:idx_leader_follower,unique"`
	FollowerID uint      `json:"follower_id" gorm:"index:idx_leader_follower,unique"`
}
