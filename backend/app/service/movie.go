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

// The general meaning of this service is to generate sql query parameter
// and get the movie list from the database using that query parameter.
func (m MovieService) GetList(qp core.QueryParams) ([]core.Movie, error) {
	movieList, err := m.movieStorage.SelectAllMovies(qp)
	if err != nil {
		return nil, fmt.Errorf("error while selecting movies: %w", err)
	}

	return movieList, nil
}

// Prepare the movie list slice for export.
func (m MovieService) GetCSV(qp core.QueryParams) ([]core.MovieCSV, error) {
	const SecondsInMinutes = 60

	movieList, err := m.movieStorage.SelectMoviesCSV(qp)
	if err != nil {
		return nil, fmt.Errorf("error while SelectMoviesCSV: %w", err)
	}

	for i := 0; i < len(movieList); i++ {
		movieList[i].Number = i + 1
		movieList[i].Duration /= SecondsInMinutes
	}

	return movieList, nil
}
