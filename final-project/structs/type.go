package structs

import "github.com/asaskevich/govalidator"

type RequestLogin struct {
	Email    string `json:"email" valid:"email~Invalid email format,required~Email is required"`
	Password string `json:"password"  valid:"required~Username is required"`
}

func (rl RequestLogin) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(rl)

	return result, err
}

type RequestUpdateUser struct {
	Username string `json:"username" valid:"required~Username is required"`
	Email    string `json:"email" valid:"email~Invalid email format,required~Email is required"`
}

func (ru RequestUpdateUser) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(ru)

	return result, err
}
