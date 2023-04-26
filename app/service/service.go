package service

import (
	"github.com/Brigant/PetPorject/config"
)

type Deps struct {
	AccountStorage  AccountStorage
	DirectorStorage DirectorStorage
	MovieStorage    MovieStorage
	ListSorage      ListSorage
}

type Services struct {
	Account  AccountService
	Director DirectorService
	Movie    MovieService
	List     ListService
}

func New(deps Deps, cfg config.Config) Services {
	return Services{
		Account:  NewAccountService(deps.AccountStorage, cfg),
		Director: NewDirectorService(deps.DirectorStorage),
		Movie:    NewMovieService(deps.MovieStorage),
		List:     NewListService(deps.ListSorage),
	}
}
