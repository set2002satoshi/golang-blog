package model

import (
	"gorm.io/gorm"
)

type CustomerInfo struct {
	gorm.Model
	Email    string   `json:"email" gorm:"unique;not null"`
	Password []byte   `json:"password" gorm:"not null"`
	Blogs    []Blog   `json:"blogs" gorm:"foreignKey:CustomerInfoID; constraint:OnDelete:CASCADE"`
	Customer Customer `gorm:"constraint:OnDelete:CASCADE"`
}

type Customer struct {
	gorm.Model
	CustomerInfoID uint
	Name           string `json:"name" gorm:"unique;not null"`
	Thumbnail      string `json:"thumbnail" gorm:"not null;size:256"`
	Message        string `json:"message" gorm:"not null;size256"`
}

type CustomerMsgForm struct {
	// Name string `json:"name"`
	// Thumbnail string `json:"thumbnail" gorm:"not null;size:256"`
	Message string `form:"message" gorm:"not null;size256"`
}

type CustomerForm struct {
	// Name string `json:"name"`
	Thumbnail string `json:"thumbnail" gorm:"not null;size:256"`
	Message   string `json:"message" gorm:"not null;size256"`
}

type CustomerInfoForm struct {
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"unique;not null"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
