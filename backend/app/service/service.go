package service

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

func New(deps Deps) Services {
	return Services{
		Account:  NewAccountService(deps.AccountStorage),
		Director: NewDirectorService(deps.DirectorStorage),
		Movie:    NewMovieService(deps.MovieStorage),
		List:     NewListService(deps.ListSorage),
	}
}
