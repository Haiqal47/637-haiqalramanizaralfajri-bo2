package models

import (
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"unique_index" json:"username" valid:"required~Username is required"`
	Email     string    `gorm:"unique" json:"email" valid:"email~Invalid email format,required~Email is required"`
	Password  string    `json:"password"  valid:"required~Username is required,minstringlength(6)~Password at least should be 6"`
	Age       uint64    `json:"age" valid:"required~Age is required,minintvalue(8)~Age need to be 8+"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user User) Validate() (bool, error) {
	govalidator.ParamTagMap["minintvalue"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		value, _ := strconv.Atoi(params[0])
		return value > 8
	})

	result, err := govalidator.ValidateStruct(user)

	return result, err
}
