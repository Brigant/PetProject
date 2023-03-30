package service

import (
	"github.com/Brigant/PetPorject/backend/app/core"
)

//go:generate mockgen -source=./contract.go -destination=./contract_mock_test.go -package=service

type AccountStorage interface {
	InsertAccount(account core.Account) (accountID string, err error)
	SelectAccountByPhone(phone string) (core.Account, error)
	SelectAccountByID(accountID string) (core.Account, error)
	InsertSession(session core.Session) (core.Session, error)
	RefreshSession(session core.Session) (core.Session, error)
}

type DirectorStorage interface {
	InsertDirector() error
	SelectDirector() error
	UpdateDirector() error
}

type MovieStorage interface {
	InsertMovie() error
	SelectAllMovies() error
	SelectMovieByID() error
}
