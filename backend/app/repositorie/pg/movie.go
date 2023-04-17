package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type MovieDB struct {
	db *sqlx.DB
}

func NewMovieDB(db *sqlx.DB) MovieDB {
	return MovieDB{db: db}
}

// Insert structure movie to database.
func (d MovieDB) InsertMovie(movie core.Movie) error {
	const expectedEffectedRow = 1

	query := `INSERT INTO public.movie(
		director_id, title, ganre, rate, release_date, duration)
		VALUES (:director_id, :title, :ganre, :rate, :release_date, :duration) RETURNING id;`

	result, err := d.db.NamedExec(query, &movie)
	if err != nil {
		pqError := new(pq.Error)
		if errors.As(err, &pqError) && pqError.Code.Name() == ErrCodeForeignKeyViolation {
			return core.ErrForeignViolation
		}

		if errors.As(err, &pqError) && pqError.Code.Name() == ErrCodeUniqueViolation {
			return core.ErrUniqueMovie
		}

		return fmt.Errorf("error in NamedEx: %w", err)
	}

	effectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("the error is in RowsAffected: %w", err)
	}

	if effectedRows != expectedEffectedRow {
		return core.ErrNowMovieAdd
	}

	return nil
}

// Select and return the movie entities via movie ID.
func (d MovieDB) SelectMovieByID(movieID string) (core.Movie, error) {
	query := `SELECT id, director_id, title, ganre, rate, release_date, duration, created, modified
	FROM public.movie WHERE id=$1`

	var movie core.Movie
	if err := d.db.Get(&movie, query, movieID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Movie{}, core.ErrMovieNotFound
		}

		return core.Movie{}, fmt.Errorf("an error occurs while getting the movie: %w", err)
	}

	return movie, nil
}

func (d MovieDB) SelectAllMovies(qp core.QueryParams) ([]core.Movie, error) {
	queryCondition := d.makeConditionQuery(qp)

	query := `SELECT id, director_id, title, genre, rate, release_date, duration, created, modified FROM public.movie `

	fullQuery := query + queryCondition

	var movieList []core.Movie
	if err := d.db.Select(&movieList, fullQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrMovieNotFound
		}

		return nil, fmt.Errorf("an error occurs while getting the movie list: %w", err)
	}

	return movieList, nil
}

func (d MovieDB) SelectMoviesCSV(qp core.QueryParams) ([]core.MovieCSV, error) {
	queryCondition := d.makeConditionQuery(qp)

	query := `SELECT m.title, m.genre, d.name as director_name, m.rate, m.release_date, m.duration FROM public.movie AS m
		INNER JOIN public.director AS d ON d.id=m.director_id `

	fullQuery := query + queryCondition

	var csvList []core.MovieCSV

	rows, err := d.db.DB.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("error while Query: %w", err)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows.Err(): %w", err)
	}

	for rows.Next() {
		var movie core.MovieCSV
		if err := rows.Scan(
			&movie.Title,
			&movie.Genre,
			&movie.DirectorName,
			&movie.Rate,
			&movie.ReleaseDate.Time,
			&movie.Duration); err != nil {
			return nil, fmt.Errorf("error while scan director list: %w", err)
		}

		csvList = append(csvList, movie)
	}

	defer rows.Close()

	return csvList, nil
}

func (d MovieDB) makeConditionQuery(queryParameter core.QueryParams) string {
	var queryCondition string

	if len(queryParameter.Filter) > 0 {
		where := "WHERE "

		for i := 0; i < len(queryParameter.Filter); i++ {
			if queryParameter.Filter[i].Val != "" {
				if _, err := strconv.Atoi(queryParameter.Filter[i].Val); err == nil {
					where = where + queryParameter.Filter[i].Key + ">=" + queryParameter.Filter[i].Val + " AND "
				} else {
					where = where + queryParameter.Filter[i].Key + "='" + queryParameter.Filter[i].Val + "' AND "
				}
			}
		}

		where = strings.TrimSuffix(where, "AND ")

		queryCondition += where
	}

	if len(queryParameter.Sort) > 0 {
		order := "ORDER BY "

		for i := 0; i < len(queryParameter.Sort); i++ {
			if queryParameter.Sort[i].Val != "" {
				order = order + queryParameter.Sort[i].Key + " " + queryParameter.Sort[i].Val + ", "
			}
		}

		order = strings.TrimSuffix(order, ", ")

		queryCondition += order
	}

	queryCondition = queryCondition + " LIMIT " + queryParameter.Limit
	queryCondition = queryCondition + " OFFSET " + queryParameter.Offset

	return queryCondition
}
