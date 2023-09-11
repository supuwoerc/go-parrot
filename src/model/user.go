package model

import (
	"go-parrot/src/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	RealName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Phone    string `gorm:"size:128"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:255;not null" json:"password,omitempty"`
}

func (user *User) EncryptPassword() error {
	password, err := utils.Encrypt(user.Password)
	if err == nil {
		user.Password = password
	}
	return err
}

// gorm创建对象之前的钩子
func (user *User) BeforeCreate(db *gorm.DB) error {
	return user.EncryptPassword()
}
