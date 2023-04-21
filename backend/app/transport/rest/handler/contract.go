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
	CreateDirector(director core.Director) error
	GetDirectorWithID(directorID string) (core.Director, error)
	GetDirectorList() ([]core.Director, error)
}

type MovieService interface {
	CreateMovie(movie core.Movie) error
	Get(movieID string) (core.Movie, error)
	GetList(core.ConditionParams) ([]core.Movie, error)
	GetCSV(core.ConditionParams) ([]core.MovieCSV, error)
}

type ListsService interface {
	Create(list core.MovieList) (string, error)
	GetAllAccountLists(core.ConditionParams) ([]core.MovieList, error)
}
