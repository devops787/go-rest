package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	Id uint64
	Name string
	Email string
	gorm.Model
}

