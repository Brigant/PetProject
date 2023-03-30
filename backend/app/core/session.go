package core

import (
	"errors"
	"time"
)

type Session struct {
	RefreshToken string    `json:"refresh_token"`
	AccountID    string    `json:"account_id"`
	Role         string    `json:"role"`
	RequestHost  string    `json:"request_host"`
	UserAgent    string    `json:"user_agent"`
	ClientIP     string    `json:"client_ip"`
	Expired      time.Time `json:"expired"`
	Created      time.Time `json:"created"`
}

var (
	ErrRefreshTokenExpired = errors.New("refresh token has expired")
	ErrSesseionNotFound    = errors.New("session is not found with such credentials")
)
