package service

import (
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
)

type MovieService struct {
	movieStorage MovieStorage
}

func NewMovieService(storage MovieStorage) MovieService {
	return MovieService{movieStorage: storage}
}

func (m MovieService) CreateMovie(movie core.Movie) (string, error) {
	if err := m.movieStorage.InsertMovie(movie); err != nil {
		return "", fmt.Errorf("error happens while inserting movie: %w", err)
	}

	return "", nil
}
