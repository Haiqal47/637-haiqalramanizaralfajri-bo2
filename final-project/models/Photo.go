package models

import "time"

type Photo struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `json:"title"  valid:"required~Title is required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" valid:"required~photo_url is required"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
