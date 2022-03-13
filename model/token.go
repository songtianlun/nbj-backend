package model

type UserRTokenModel struct {
	BaseModel
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	UserUUID     string `json:"user_uuid" gorm:"column:user_uuid"`
}

type UserATokenModel struct {
	BaseModel
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	RefreshUUID string `json:"refresh_uuid" gorm:"column:refresh_uuid"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
