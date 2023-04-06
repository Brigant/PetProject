package pg

import (
	"errors"
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DirectorDB struct {
	db *sqlx.DB
}

func NewDirectorDB(db *sqlx.DB) DirectorDB {
	return DirectorDB{db: db}
}

// The method inserts the director to the DB.
func (d DirectorDB) InsertDirector(director core.Director) error {
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

	if affectedRow != 1 {
		return errors.New("wrong number of affected rows , should be 1")
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
		return core.Director{}, fmt.Errorf("errow while Select director: %w", err)
	}

	return director, nil
}

// The method grabs the all directors and returns it in the slice.
func (d DirectorDB) SelectDirectorList() ([]core.Director, error) {
	query := `SELECT id, name, birth_date::timestamp, created, modified
		FROM public.director`

	var directorsList []core.Director

	rows, _ := d.db.DB.Query(query)
	defer rows.Close()
	for rows.Next() {
		var director core.Director
		if err := rows.Scan(
			&director.ID,
			&director.Name,
			&director.BirthDate.Time,
			&director.Created,
			&director.Modified); err != nil {
			return nil, err
		}

		directorsList = append(directorsList, director)
	}

	return directorsList, nil
}
