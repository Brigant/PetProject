package pg

import (
	"errors"
	"fmt"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ListDB struct {
	db *sqlx.DB
}

func NewListDB(db *sqlx.DB) ListDB {
	return ListDB{db: db}
}

func (d ListDB) Insert(list core.MovieList) (string, error) {
	query := `INSERT INTO public.list(type, account_id, movie_id)
		VALUES(:type, :account_id, :movie_id) RETURNING id`

	rows, err := d.db.NamedQuery(query, &list)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", core.ErrDuplicateRow
		}

		return "", fmt.Errorf("insterting error: %w", err)
	}
	defer rows.Close()

	var listID string

	rows.Next()

	if err := rows.Scan(&listID); err != nil {
		return "", fmt.Errorf("error while scaning: %w", err)
	}

	return listID, nil
}

func (d ListDB) SelectMovieListByAccountID(accountID string, queryParam core.ConditionParams) ([]core.MovieList, error) {
	var list []core.MovieList

	query := ``

	condition := buildQueryCondition(queryParam)

	query += condition

	rows, err := d.db.NamedQuery(query, &list)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return nil, core.ErrDuplicateRow
		}

		return nil, fmt.Errorf("insterting error: %w", err)
	}
	defer rows.Close()

	var listID string

	rows.Next()

	if err := rows.Scan(&listID); err != nil {
		return nil, fmt.Errorf("error while scaning: %w", err)
	}

	return list, nil
}
