package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Ip string
}
