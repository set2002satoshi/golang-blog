package model

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	CustomerInfoID uint
	UserName       string
	BlogImage      string `json:"blog_image" gorm:"size:256"`
	Title          string `json:"title" gorm:"size:25;not null"`
	Subtitle       string `json:"subtitle" gorm:"size:100;not null"`
	Content        string `json:"content"  gorm:"size:265;not null"`
	Tags           []Tag  `json:"tags" gorm:"many2many:blog_tags; constraint:OnDelete:CASCADE"`
}

type BlogForm struct {
	// BlogImage  string `form:"blog_image" gorm:"size:256"`
	Title    string `form:"title" gorm:"size:25;not null"`
	Subtitle string `form:"subtitle" gorm:"size:100;not null"`
	Content  string `form:"content"  gorm:"size:265;not null"`
	Tag      int    `form:"tag"`
}
