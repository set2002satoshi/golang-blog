package model

import (
	"gorm.io/gorm"
)

type CustomerInfo struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"password" gorm:"not null"`
	Blogs    []Blog `json:"blogs" gorm:"foreignKey:CustomerInfoID"`
	Customer Customer
}

type Customer struct {
	gorm.Model
	CustomerInfoID uint
	Name           string `json:"name" gorm:"not null"`
	Thumbnail      string `json:"thumbnail" gorm:"not null;size:256"`
	Message        string `json:"message" gorm:"not null;size256"`
}

type CustomerForm struct {
	Name string `json:"name"`
	// Thumbnail string `json:"thumbnail" gorm:"not null;size:256"`
	Message string `json:"message" gorm:"not null;size256"`
}

type CustomerInfoForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
