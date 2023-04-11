package pg

import (
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
		return fmt.Errorf("error in RowsAffected: %w", err)
	}

	if effectedRows != expectedEffectedRow {
		return core.ErrNowMovieAdded
	}

	return nil
}

func (d MovieDB) SelectAllMovies() error {
	return nil
}

func (d MovieDB) SelectMovieByID() error {
	return nil
}
