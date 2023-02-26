package service

type MovieService struct {
	storage MovieStorage
}

func NewMovieService(storage MovieStorage) MovieService {
	return MovieService{storage: storage}
}

// TODO: implement the func.
func (m MovieService) CreateMovie() error {
	return nil
}
