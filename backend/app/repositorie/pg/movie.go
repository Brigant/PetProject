package pg

import "github.com/jmoiron/sqlx"

type MovieDB struct {
	db *sqlx.DB
}

func NewMovieDB(db *sqlx.DB) MovieDB {
	return MovieDB{db: db}
}

func (d MovieDB) InsertMovie() error {
	return nil
}

func (d MovieDB) SelectAllMovies() error {
	return nil
}

func (d MovieDB) SelectMovieByID() error {
	return nil
}
