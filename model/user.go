package model

import "gorm.io/gorm"

// User 数据库结构体
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(12);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
