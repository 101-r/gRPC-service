package model

import (
	"time"
)

type User struct {
	Id       int64
	UserInfo UserInfo
}

type UserInfo struct {
	Id        int64
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
