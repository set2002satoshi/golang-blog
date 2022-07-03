package model

import (
	"time"
)

type ProfilePage struct {
	UserName      string
	UserThumbnail string
	Message       string
	Profile       []ProfileBlog
	CreatedAt     time.Time
}

type ProfileBlog struct {
	BlogID        string
	BlogThumbnail string
	Title         string
}
