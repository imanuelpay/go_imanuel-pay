package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}
