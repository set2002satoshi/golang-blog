package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	BlogID int
	Tag    string `json:"tag" grom:"not null"`
}

