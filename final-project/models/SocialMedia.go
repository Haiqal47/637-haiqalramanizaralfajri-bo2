package models

type SocialMedia struct {
	ID             uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name           string `json:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"social_media_url~Username is required"`
	UserId         uint64 `json:"user_id"`
}
