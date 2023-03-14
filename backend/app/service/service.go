package service

type Deps struct {
	AccountStorage  AccountStorage
	DirectorStorage DirectorStorage
	MovieStorage    MovieStorage
}

type Services struct {
	Account  AccountService
	Director DirectorService
	Movie    MovieService
}

func New(deps Deps) Services {
	return Services{
		Account:  NewAccountService(deps.AccountStorage),
		Director: NewDirectorService(deps.DirectorStorage),
		Movie:    NewMovieService(deps.MovieStorage),
	}
}
