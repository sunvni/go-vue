package entity

import (
	"github.com/jinzhu/gorm"
)

type Tweet struct {
    gorm.Model
	Title string
    Content string
}