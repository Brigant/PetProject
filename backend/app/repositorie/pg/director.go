package pg

import "github.com/jmoiron/sqlx"

type DirectorDB struct {
	db *sqlx.DB
}

func NewDirectorDb(db *sqlx.DB) DirectorDB {
	return DirectorDB{db: db}
}

// TODO: implement the func.
func (d DirectorDB) InsertDirector() error {
	return nil
}

// TODO: implement the func.
func (d DirectorDB) SelectDirector() error {
	return nil
}

// TODO: implement the func.
func (d DirectorDB) UpdateDirector() error {
	return nil
}
