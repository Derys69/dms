package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"uniqueIndex"`
	Password   string
	GoogleID   string
	Role       string `gorm:"default:'pengguna_umum'"`
	Department string
}
