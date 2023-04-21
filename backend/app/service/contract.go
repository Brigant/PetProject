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
	SelectSession(session core.Session) (core.Session, error)
	RefreshSession(session core.Session) error
	DeleteSesions(accountID string) error
}

type DirectorStorage interface {
	InsertDirector(director core.Director) error
	SelectDirectorByID(directorID string) (core.Director, error)
	SelectDirectorList() ([]core.Director, error)
}

type MovieStorage interface {
	InsertMovie(movie core.Movie) error
	SelectAllMovies(core.ConditionParams) ([]core.Movie, error)
	SelectMoviesCSV(core.ConditionParams) ([]core.MovieCSV, error)
	SelectMovieByID(movieID string) (core.Movie, error)
}

type ListSorage interface {
	Insert(core.MovieList) (string, error)
	SelectAllUsersLists(core.ConditionParams) ([]core.MovieList, error)
}
