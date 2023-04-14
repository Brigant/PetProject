package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // the blank import is needed beceause of sqlx requirements
)

type DirectorDB struct {
	db *sqlx.DB
}

func NewDirectorDB(db *sqlx.DB) DirectorDB {
	return DirectorDB{db: db}
}

// The method inserts the director to the DB.
func (d DirectorDB) InsertDirector(director core.Director) error {
	const expectedEffectedRow = 1

	query := `INSERT INTO public.director(name, birth_date)
		VALUES($1, $2)`

	result, err := d.db.DB.Exec(query, director.Name, director.BirthDate.Time)
	if err != nil {
		return fmt.Errorf("can't exec because: %w", err)
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("err happaned while getin RowsAffected: %w", err)
	}

	if affectedRow != expectedEffectedRow {
		return core.ErrNowDirectorAdded
	}

	return nil
}

// The method selects the director speciofied by ID and returns it.
func (d DirectorDB) SelectDirectorByID(directorID string) (core.Director, error) {
	query := `SELECT id, name, birth_date, created, modified
		FROM public.director
		WHERE id=$1`

	var director core.Director

	err := d.db.DB.QueryRow(query, directorID).Scan(
		&director.ID, &director.Name, &director.BirthDate.Time, &director.Created, &director.Modified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Director{}, core.ErrNowDirectorFound
		}
		
		return core.Director{}, fmt.Errorf("errow while Select director: %w", err)
	}

	return director, nil
}

// The method grabs the all directors and returns it in the slice.
func (d DirectorDB) SelectDirectorList() ([]core.Director, error) {
	query := `SELECT id, name, birth_date::timestamp, created, modified
		FROM public.director`

	var directorsList []core.Director

	rows, err := d.db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while Query: %w", err)
	}

	defer rows.Close()

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows.Err(): %w", err)
	}

	for rows.Next() {
		var director core.Director
		if err := rows.Scan(
			&director.ID,
			&director.Name,
			&director.BirthDate.Time,
			&director.Created,
			&director.Modified); err != nil {
			return nil, fmt.Errorf("error while scan director list: %w", err)
		}

		directorsList = append(directorsList, director)
	}

	return directorsList, nil
}
