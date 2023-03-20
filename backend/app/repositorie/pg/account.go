package pg

import (
	"errors"
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AccountDB struct {
	db *sqlx.DB
}

func NewAccountDB(db *sqlx.DB) AccountDB {
	return AccountDB{db: db}
}

const (
	ErrCodeUniqueViolation = "unique_violation"
	ErrCodeNoData          = "no_data"
)

// Insert the account model to databese and returning the newly created account id.
func (r AccountDB) InsertAccount(account core.Account) (string, error) {
	var id string

	query := "INSERT INTO public.account(phone, password, age, role) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := r.db.DB.QueryRow(query, account.Phone, account.Password, account.Age, account.Role).Scan(&id); err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", core.ErrDuplicatePhone
		}

		return "", fmt.Errorf("cannot execute query: %w", err)
	}

	return id, nil
}

func (r AccountDB) SelectAccount() error {
	return nil
}
