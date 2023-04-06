package handler

import "github.com/Brigant/PetPorject/backend/app/core"

//go:generate mockgen -source=./contract.go -destination=./contract_mock_test.go -package=handler

type AccountService interface {
	CreateUser(account core.Account) (id string, err error)
	Login(login, password string, session core.Session) (core.TokenPair, error)
	ParseToken(string) (string, string, error)
	RefreshTokenpair(session core.Session) (core.TokenPair, error)
	Logout(accountID string) error
}

type DirectorService interface {
	CreatDirector(director core.Director) error
	GetDirectorWithName(directorName string) (core.Director, error)
	GetDirectorWithID(directorID string) (core.Director, error)
	GetDirectorList() ([]core.Director, error)
	UpdateDirector() error
}

type MovieService interface {
	GetMovie() error
}

type ListsService interface {
	GetList() error
}
