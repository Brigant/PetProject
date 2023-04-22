package pg

import (
	"errors"
	"fmt"
	"strings"

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

func (d ListDB) SelectAllUsersLists(conditions []core.QuerySliceElement) ([]core.MovieList, error) {
	var list []core.MovieList

	query := `SELECT * FROM public.list `

	where := d.prepareCond(conditions)

	fullQuery := query + where

	fmt.Println(fullQuery)

	if err := d.db.Select(&list, fullQuery); err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUndefinedColumn {
			return nil, core.ErrUnkownConditionKey
		}

		return nil, fmt.Errorf("an error occurs while getting the movie list: %w", err)
	}

	return list, nil
}

func (d ListDB) prepareCond(cond []core.QuerySliceElement) string {
	where := "WHERE "

	if len(cond) == 1 {
		where += cond[0].Key + "='" + cond[0].Val + "'"

		return where
	}

	for i := 0; i < len(cond); i++ {
		if i == 0 {
			where += cond[0].Key + "='" + cond[0].Val + "' AND ("
		} else {
			where = where + cond[i].Key + "='" + cond[i].Val + "' OR "
		}
	}

	where = strings.TrimSuffix(where, " OR ") + ")"

	return where
}
