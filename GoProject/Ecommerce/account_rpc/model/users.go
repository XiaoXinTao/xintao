package model

import (
	"time"
)

func (Users) TableName() string {
	return "users"
}

func NewUsers() *Users {
	return &Users{}
}

type Users struct {
	PassportUid  int64     `gorm:"column:passport_uid" json:"passport_uid"`   // 用户账号
	PasswordCode string    `gorm:"column:password_code" json:"password_code"` // 用户密码
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`     // create_time
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`     // update_time
}