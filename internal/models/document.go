package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Title      string
	Content    string
	OwnerID    uint
	Department string
}
