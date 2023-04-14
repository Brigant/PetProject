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

// Add the movie to the storage.
func (m MovieService) CreateMovie(movie core.Movie) error {
	err := m.movieStorage.InsertMovie(movie)
	if err != nil {
		return fmt.Errorf("error happens while inserting movie: %w", err)
	}

	return nil
}

// The simple get the movie from the storage.
func (m MovieService) Get(movieID string) (core.Movie, error) {
	movie, err := m.movieStorage.SelectMovieByID(movieID)
	if err != nil {
		return core.Movie{}, fmt.Errorf("service Get got the error: %w", err)
	}

	return movie, nil
}
