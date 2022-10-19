package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username        string    `gorm:"uniqueIndex" json:"username" valid:"required~Username is required" form:"username"`
	Email           string    `gorm:"uniqueIndex" json:"email" valid:"email~Invalid email format,required~Email is required" form:"email"`
	Password        string    `json:"password"  valid:"required~Username is required,minstringlength(6)~Password at least should be 6" form:"password"`
	Age             uint64    `json:"age" valid:"required~Age is required" form:"age"`
	ProfileImageUrl string    `json:"profile_image_url" valid:"required~profile_image_url is required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (user User) Validate() (bool, error) {
	if user.Age < 8 {
		return false, errors.New("age need to be 8+")
	}

	result, err := govalidator.ValidateStruct(user)

	return result, err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var user User

	errModel := tx.Model(&User{}).Where("username = ?", u.Username).First(&user).Error
	if errModel == nil {
		err = errors.New("username already used")
	}

	errModel = tx.Model(&User{}).Where("email = ?", u.Email).First(&user).Error
	if errModel == nil {
		err = errors.New("email already used")
	}

	return
}
