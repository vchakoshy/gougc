package models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"column:username;type:varchar(255);uniqueIndex"`
	Password  string `gorm:"column:password;type:varchar(255)"`
	Email     string `gorm:"column:email;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
