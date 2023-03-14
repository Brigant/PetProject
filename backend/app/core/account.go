package core

import (
	"errors"
)

type Account struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" binding:"required,e164,lowercase"`
	Password string `json:"password" binding:"required,min=8,max=255,ascii"`
	Age      int    `json:"age" binding:"required,gte=1,lte=120"`
	Role     string `jsdon:"role" binding:"required,max=255,lowercase"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

var (
	ErrDuplicatePhone = errors.New("phone is already exists in the database")
	ErrUserNotFound   = errors.New("user is not found with such credentials")
)
