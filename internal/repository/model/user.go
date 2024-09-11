package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id       int
	UserInfo UserInfo
}

type UserInfo struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
