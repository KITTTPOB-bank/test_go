package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Hospital  string `json:"hospital"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Hospital string `json:"hospital"`
}
