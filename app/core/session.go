package core

import (
	"errors"
	"time"
)

type Session struct {
	ID        string        `json:"id"`
	AccountID string        `json:"account_id"`
	ExpiredIn time.Duration `json:"expired_in" postgres:"expired_in"`
	Created   time.Time     `json:"created"`
}

var (
	ErrRefreshTokenExpired = errors.New("refresh token has expired")
	ErrSesseionNotFound    = errors.New("session is not found with such credentials")
)
