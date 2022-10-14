package models

import "time"

type Comment struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserId    uint64    `json:"user_id"`
	PhotoId   uint64    `json:"photo_id"`
	Message   string    `json:"message" valid:"required~Message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
