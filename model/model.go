package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"column:id;AUTO_INCREMENT;comment:序号;unique" `
	UUID      string     `gorm:"primary_key;not null;column:uuid;unique" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type Token struct {
	Token string `json:"token"`
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	bm.UUID = uuid.New().String()
	return
}
