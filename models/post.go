package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	gorm.Model
	UserId int `gorm:"index:user_id"`
	Text string `sql:"type:text;"`
}

