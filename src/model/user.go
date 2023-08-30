package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	RealName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Phone    string `gorm:"size:128"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:255;not null" json:"password,omitempty"`
}
