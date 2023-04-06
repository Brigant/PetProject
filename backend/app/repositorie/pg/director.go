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

func (d DirectorDB) InsertDirector(director core.Director) error {
	query := `INSERT INTO public.director(name, birth_date)
		VALUES($1, $2)`

	result, err := d.db.DB.Exec(query, director.Name, director.BirthDate)
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

func (d DirectorDB) SelectDirectorByName(directorName string) (core.Director, error) {
	query := `SELECT id, name, birth_date, created, modified
		FROM public.director
		WHERE name=$1`

	var director core.Director

	err := d.db.DB.QueryRow(query, directorName).Scan(
		&director.ID, &director.Name, &director.BirthDate, &director.Created, &director.Modified)
	if err != nil {
		return core.Director{}, fmt.Errorf("errow while Select director: %w", err)
	}

	return director, nil
}

func (d DirectorDB) SelectDirectorByID(directorID string) (core.Director, error) {
	query := `SELECT id, name, birth_date, created, modified
		FROM public.director
		WHERE id=$1`

	var director core.Director

	err := d.db.DB.QueryRow(query, directorID).Scan(
		&director.ID, &director.Name, &director.BirthDate, &director.Created, &director.Modified)
	if err != nil {
		return core.Director{}, fmt.Errorf("errow while Select director: %w", err)
	}

	return director, nil
}

func (d DirectorDB) SelectDirectorList() ([]core.Director, error) {
	query := `SELECT id, name, birth_date, created, modified
		FROM public.director`

	var directorsList []core.Director

	if err := d.db.Select(&directorsList, query); err != nil {
		return nil, fmt.Errorf("error while Select director list: %w", err)
	}

	return directorsList, nil
}

func (d DirectorDB) UpdateDirector() error {
	return nil
}
