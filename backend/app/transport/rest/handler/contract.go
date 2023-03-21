package handler

import "github.com/Brigant/PetPorject/backend/app/core"

//go:generate mockgen -source=./contract.go -destination=./mocks/contract_mock_test.go -package=handler

type AccountService interface {
	CreateUser(account core.Account) (id string, err error)
	GetUser() error
	ParseToken(string) (string, string, error)
}

type DirectorService interface {
	CreatDirector() error
	GetDirector() error
	UpdateDirector() error
}

type MovieService interface {
	GetMovie() error
}

type ListsService interface {
	GetList() error
}
