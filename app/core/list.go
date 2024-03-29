package core

import (
	"errors"

	"github.com/google/uuid"
)

type MovieList struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Type      string    `json:"type" binding:"required" db:"type"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	Created   string    `json:"created" db:"created"`
	Modified  string    `json:"modified" db:"modified"`
}

var (
	ErrDuplicateRow        = errors.New("such record already exists")
	ErrEmptyMovieListType  = errors.New("the movie list type should't be empty")
	ErrEpmtryMovieID       = errors.New("the movie ID should't be empty")
	ErrForeignKeyViolation = errors.New("some value has no reference to the list or to the movie")
)
