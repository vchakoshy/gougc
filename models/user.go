package models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Username  string `gorm:"column:username;type:varchar(255);uniqueIndex" json:"username"`
	Password  string `gorm:"column:password;type:varchar(255)" json:"-"`
	Email     string `gorm:"column:email;type:varchar(255)" json:"-"`
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
