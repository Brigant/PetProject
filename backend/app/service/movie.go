package service

import (
	"fmt"
	"strconv"
	"strings"

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
	var queryParams string

	if len(qp.Filter) > 0 {
		where := "WHERE "

		for field, value := range qp.Filter {
			if value != "" {
				if _, err := strconv.Atoi(value); err == nil {
					where = where + field + ">" + value + " AND "
				} else {
					where = where + field + "='" + value + "' AND "
				}
			}
		}

		where = strings.TrimSuffix(where, "AND ")

		queryParams = queryParams + where
	}

	if len(qp.Sort) > 0 {
		order := "ORDER BY "

		for field, value := range qp.Sort {
			if value != "" {
				order = order + field + " " + value + ", "
			}
		}

		order = strings.TrimSuffix(order, ", ")

		queryParams = queryParams + order
	}

	queryParams = queryParams + " LIMIT " + qp.Limit
	queryParams = queryParams + " OFFSET " + qp.Offset

	movieList, err := m.movieStorage.SelectAllMovies(queryParams)
	if err != nil {
		return nil, fmt.Errorf("service Get got the error: %w", err)
	}

	return movieList, nil
}
