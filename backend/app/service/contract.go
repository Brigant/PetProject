package service

type AccountStorage interface {
	InsertAccount() error
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
