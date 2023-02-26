package handler

type AccountService interface {
	CreateUser() error
	GetUser() error
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