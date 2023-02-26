package pg

import (
	"github.com/jmoiron/sqlx"
)

type AccountDb struct {
	db *sqlx.DB
}

func NewAccountDb(db *sqlx.DB) AccountDb {
	return AccountDb{db: db}
}

// TODO: implement the func.
func (r AccountDb) InsertAccount() error {
	return nil
}

// TODO: implement the func.
func (r AccountDb) SelectAccount() error {
	return nil
}
