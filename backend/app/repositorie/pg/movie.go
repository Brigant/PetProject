package pg

import (
	"database/sql"
	"errors"
	"fmt"

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

func (d MovieDB) SelectAllMovies() error {
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
