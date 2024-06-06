package models

import (
	"database/sql"
	"time"
)

type ContentType string

const (
	Image ContentType = "IMAGE"
	Video ContentType = "VIDEO"
	Audio ContentType = "AUDIO"
)

type Attachment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	ContentType ContentType    `gorm:"column:content_type;type:varchar(8)" json:"content_type"`
	Size        uint           `json:"size"`
	Duration    sql.NullString `gorm:"column:duration;type:varchar(10)" json:"duration"`
	FileURL     string         `gorm:"column:file_url;type:varchar(500)" json:"file_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
