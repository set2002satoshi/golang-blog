package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Blogs []Blog `json:"blogs" gorm:"many2many:blog_tags;"`
	Tag  string  `json:"tag"`
}
