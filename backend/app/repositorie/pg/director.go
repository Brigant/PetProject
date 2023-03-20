package pg

import "github.com/jmoiron/sqlx"

type DirectorDB struct {
	db *sqlx.DB
}

func NewDirectorDB(db *sqlx.DB) DirectorDB {
	return DirectorDB{db: db}
}

func (d DirectorDB) InsertDirector() error {
	return nil
}

func (d DirectorDB) SelectDirector() error {
	return nil
}

func (d DirectorDB) UpdateDirector() error {
	return nil
}
