package core

import (
	"errors"
	"time"
)

type Session struct {
	RefreshToken string        `json:"refresh_token"`
	AccountID    string        `json:"account_id"`
	RequestHost  string        `json:"request_host"`
	UserAgent    string        `json:"user_agent"`
	ClientIP     string        `json:"client_ip"`
	ExpiredIn    time.Duration `json:"expired_in" postgres:"expired_in"`
	Created      time.Time     `json:"created"`
}

var (
	ErrRefreshTokenExpired = errors.New("refresh token has expired")
	ErrSesseionNotFound    = errors.New("session is not found with such credentials")
)
