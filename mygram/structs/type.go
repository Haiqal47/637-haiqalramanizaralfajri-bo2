package structs

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RequestLogin struct {
	Email    string `json:"email" valid:"email~Invalid email format,required~Email is required"`
	Password string `json:"password"  valid:"required~Username is required"`
}

func (rl RequestLogin) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(rl)

	return result, err
}

type RequestUpdateUser struct {
	Username        string `json:"username" valid:"required~Username is required" form:"username"`
	Email           string `json:"email" valid:"email~Invalid email format,required~Email is required" form:"email"`
	ProfileImageUrl string `json:"profile_image_url" form:"profile_image_url"`
}

func (ru RequestUpdateUser) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(ru)

	return result, err
}

type ResponseUserRegister struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint64 `json:"age"`
}
type ResponseUserLogin struct {
	Token string `json:"token"`
}

type ResponseUserUpdate struct {
	ID              uint64    `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Age             uint64    `json:"age"`
	ProfileImageUrl string    `json:"profile_image_url"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseCreatePhoto struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseUpdatePhoto struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponsePhotos struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint64 `json:"user_id"`
	User     struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"User"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestCreateComment struct {
	PhotoId uint64 `json:"photo_id"`
	Message string `json:"message"`
}

type RequestUpdateComment struct {
	Message string `json:"message"`
}

type ResponseCreateComment struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	PhotoId   uint64    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseUpdateComment struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	PhotoId   uint64    `json:"photo_id"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseComments struct {
	ID      uint64 `json:"id"`
	UserId  uint64 `json:"user_id"`
	PhotoId uint64 `json:"photo_id"`
	Message string `json:"message"`
	User    struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"User"`
	Photo struct {
		ID       uint64 `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserId   uint64 `json:"user_id"`
	} `json:"Photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestCreateSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type RequestUpdateSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type ResponseCreateSocialMedia struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint64    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ResponseUpdateSocialMedia struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint64    `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ResponseGetSocialMedia struct {
	SocialMedias []ResponseSocialMedia `json:"social_medias"`
}

type ResponseSocialMedia struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint64 `json:"user_id"`
	User           struct {
		ID              uint64 `json:"id"`
		Username        string `json:"username"`
		ProfileImageUrl string `json:"profile_image_url"`
	} `json:"User"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
