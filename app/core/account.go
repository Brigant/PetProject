package core

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Account struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" binding:"required,e164,lowercase"`
	Password string `json:"password" binding:"required,min=8,max=255,ascii"`
	Age      int    `json:"age" binding:"required,gte=1,lte=120"`
	Role     string `jsdon:"role" binding:"required,lowercase,checkRole"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

var (
	ErrDuplicatePhone        = errors.New("phone is already exists in the database")
	ErrUserNotFound          = errors.New("user is not found with such credentials")
	ErrWrongPassword         = errors.New("wrong passord")
	ErrContexAccountNotFound = errors.New("no account found in contex")
)

func SHA256(password, salt string) string {
	sum := sha256.Sum256([]byte(password + salt))
	fmt.Println("---------------------*************", salt)

	return fmt.Sprintf("%x", sum)
}
