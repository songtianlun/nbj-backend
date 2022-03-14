package model

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	UserID       uint64 `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//type UserATokenModel struct {
//	BaseModel
//	AccessToken string    `json:"access_token" gorm:"column:access_token"`
//	UserUID     uint64    `json:"user_uid" gorm:"column:user_uid"`
//	ExpiresAt   time.Time `json:"exp" gorm:"column:exp"`
//	IssuedAt    time.Time `json:"iat" gorm:"column:iat"`
//	NotBefore   time.Time `json:"nbf" gorm:"column:nbf"`
//}

type UserRTokenModel struct {
	BaseModel
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token"`
	UserUID      uint64    `json:"user_uid" gorm:"column:user_uid"`
	UserAddr     string    `json:"user_addr" gorm:"column:user_addr"`
	UserName     string    `json:"user_name" gorm:"column:user_name"`
	ExpiresAt    time.Time `json:"exp" gorm:"column:exp"`
	IssuedAt     time.Time `json:"iat" gorm:"column:iat"`
	NotBefore    time.Time `json:"nbf" gorm:"column:nbf"`
}

//func (u *UserATokenModel) RegisterAccessToken() error {
//	return DB.DB.Create(&u).Error
//}
//
//func (u *UserATokenModel) LogoutAccessToken() error {
//	return DB.DB.Delete(&u).Error
//}

func (u *UserRTokenModel) RegisterRefreshToken() error {
	return DB.DB.Create(&u).Error
}

func (u *UserRTokenModel) LogoutRefreshToken() error {
	return DB.DB.Delete(&u).Error
}

func LogoutRefreshTokenWithToken(rt string) error {
	return DB.DB.Where("refresh_token = ?", rt).Delete(&UserRTokenModel{}).Error
}

//func (u *UserATokenModel) BeforeSave(tx *gorm.DB) (err error) {
//	// 将已过期的 rt 全部注销
//	go func() { DB.DB.Where("exp <= ?", time.Now()).Delete(&UserATokenModel{}) }()
//
//	return
//}

func (u *UserRTokenModel) BeforeSave(tx *gorm.DB) (err error) {
	//var RTokens []UserRTokenModel
	// 将当前 IP 地址下的 rt 全部注销
	//DB.DB.Where("user_uid = ? AND user_addr = ?", u.UserUID, u.UserAddr).Find(&RTokens)
	//for _, v := range RTokens {
	//	go func(um UserRTokenModel) {
	//		if err := um.LogoutRefreshToken(); err != nil {
	//			logger.DebugF("error of Logout refresh token [%s] with %s",
	//				um.RefreshToken, err.Error())
	//		}
	//	}(v)
	//}

	// 将已过期的 rt 全部注销
	go func() { DB.DB.Where("exp <= ?", time.Now()).Delete(&UserRTokenModel{}) }()

	return
}

func ReTokenEffective(rt string) (*UserRTokenModel, error) {
	var d *gorm.DB
	urt := &UserRTokenModel{}
	d = DB.DB.Where("refresh_token = ?", rt).First(&urt)

	return urt, d.Error
}

//func ClearUpRefreshTokenModel(u *UserRTokenModel) (err error) {
//	DB.DB.Where("user_uuid = ? AND user_addr = ?", u.UserUID, u.UserAddr)
//}
