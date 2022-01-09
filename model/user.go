package model

import (
	"github.com/go-playground/validator/v10"
	"minepin-backend/pkg/auth"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Email    string `json:"email" gorm:"column:email" validate:"max=64"`
	Phone    string `json:"phone" gorm:"column:phone" validate:"max=16"`
}

func (u *UserModel) Create() error {
	return DB.DB.Create(&u).Error
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate 数据类型校验
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
