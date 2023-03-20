package service

type MovieService struct {
	storage MovieStorage
}

func NewMovieService(storage MovieStorage) MovieService {
	return MovieService{storage: storage}
}

func (m MovieService) CreateMovie() error {
	return nil
}
