package core

import (
	"errors"

	"github.com/google/uuid"
)

type MovieList struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	MovieID   uuid.UUID `json:"movie_id" db:"movie_id"`
	Created   string    `json:"created" db:"created"`
	Modified  string    `json:"modified" db:"modified"`
}

var ErrDuplicateRow = errors.New("the account has that list type")
