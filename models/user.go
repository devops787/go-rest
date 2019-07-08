package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"size:32" json:"name"`
	Email string `gorm:"size:255" json:"email"`
}

