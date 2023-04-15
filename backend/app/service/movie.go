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
	var queryCondition string

	if len(qp.Filter) > 0 {
		where := "WHERE "

		for i := 0; i < len(qp.Filter); i++ {
			if qp.Filter[i].Val != "" {
				if _, err := strconv.Atoi(qp.Filter[i].Val); err == nil {
					where = where + qp.Filter[i].Key + ">=" + qp.Filter[i].Val + " AND "
				} else {
					where = where + qp.Filter[i].Key + "='" + qp.Filter[i].Val + "' AND "
				}
			}
		}

		where = strings.TrimSuffix(where, "AND ")

		queryCondition = queryCondition + where
	}

	if len(qp.Sort) > 0 {
		order := "ORDER BY "

		for i := 0; i < len(qp.Sort); i++ {
			if qp.Sort[i].Val != "" {
				order = order + qp.Sort[i].Key + " " + qp.Sort[i].Val + ", "
			}
		}

		order = strings.TrimSuffix(order, ", ")

		queryCondition = queryCondition + order
	}

	queryCondition = queryCondition + " LIMIT " + qp.Limit
	queryCondition = queryCondition + " OFFSET " + qp.Offset

	movieList, err := m.movieStorage.SelectAllMovies(queryCondition)
	if err != nil {
		return nil, fmt.Errorf("error while selecting movies: %w", err)
	}

	return movieList, nil
}
