package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `json:"name" valid:"required~Name is required"`
	SocialMediaUrl string    `json:"social_media_url" valid:"required~social_media_url is required"`
	UserId         uint64    `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"User" valid:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (socialMedia SocialMedia) Validate() (bool, error) {

	result, err := govalidator.ValidateStruct(socialMedia)
	return result, err
}
