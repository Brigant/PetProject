package service

import "github.com/Brigant/PetPorject/backend/app/core"

type AccountStorage interface {
	InsertAccount(account core.Account) (id string, err error)
	SelectAccount() error
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
