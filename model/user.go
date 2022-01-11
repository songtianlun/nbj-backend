package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"minepin-backend/pkg/auth"
	"minepin-backend/pkg/logger"
	"regexp"
)

type UserModel struct {
	BaseModel
	Email    string `json:"email" gorm:"column:email;unique" validate:"max=64"`
	Phone    string `json:"phone" gorm:"column:phone;unique" validate:"max=16"`
	Password string `json:"password" gorm:"column:password;not null" validate:"min=5,max=128"`
	Nickname string `json:"nickname" gorm:"column:nickname;not null;default:-" validate:"min=1,max=32"`
}

type UserBind struct {
	Username string `json:"username" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" binding:"required" validate:"min=5,max=128"`
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

// GetUser 根据邮箱/手机号获取用户信息
func GetUser(username string) (*UserModel, error) {
	var d *gorm.DB
	u := &UserModel{}
	reg := regexp.MustCompile(
		`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)
	if reg.MatchString(username) {
		logger.Debug("Login with email: "+username)
		d = DB.DB.Where("email = ?", username).First(&u)
	} else {
		logger.Debug("Login with phone: "+username)
		d = DB.DB.Where("phone = ?", username).First(&u)
	}

	return u, d.Error
}
