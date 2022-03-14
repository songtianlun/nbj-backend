package model

type UserLoginLog struct {
	BaseModel
	UserUID      uint64 `json:"user_uid" gorm:"column:user_uid"`
	AccessToken  string `json:"access_token" gorm:"column:access_token"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	UserAddr     string `json:"user_addr" gorm:"column:user_addr"`
	UserName     string `json:"user_name" gorm:"column:user_name"`
}

func (u *UserLoginLog) Create() error {
	return DB.DB.Create(&u).Error
}
