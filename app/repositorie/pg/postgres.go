package pg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Brigant/PetPorject/app/core"
	"github.com/Brigant/PetPorject/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nececarry blank import
)

const (
	ErrCodeUniqueViolation     = "unique_violation"
	ErrCodeNoData              = "no_data"
	ErrCodeForeignKeyViolation = "foreign_key_violation"
	ErrCodeUndefinedColumn     = "undefined_column"
)

type Repository struct {
	AccountDB  AccountDB
	DirectorDB DirectorDB
	MovieDB    MovieDB
	ListDB     ListDB
}

// NewPostgresDB function returns object of datatabase.
func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	database, err := sqlx.Connect("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Database, cfg.DB.Password, cfg.DB.SSLmode))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}

	return database, nil
}

// Returns an object of the Ropository.
func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		AccountDB:  NewAccountDB(db),
		DirectorDB: NewDirectorDB(db),
		MovieDB:    NewMovieDB(db),
		ListDB:     NewListDB(db),
	}
}

func buildQueryCondition(condiotion core.ConditionParams) string {
	var queryCondition string

	if len(condiotion.Filter) > 0 {
		where := "WHERE "

		for i := 0; i < len(condiotion.Filter); i++ {
			if condiotion.Filter[i].Val != "" {
				if _, err := strconv.Atoi(condiotion.Filter[i].Val); err == nil {
					where = where + condiotion.Filter[i].Key + ">=" + condiotion.Filter[i].Val + " AND "
				} else {
					where = where + condiotion.Filter[i].Key + "='" + condiotion.Filter[i].Val + "' AND "
				}
			}
		}

		where = strings.TrimSuffix(where, "AND ")

		queryCondition += where
	}

	if len(condiotion.Sort) > 0 {
		order := "ORDER BY "

		for i := 0; i < len(condiotion.Sort); i++ {
			if condiotion.Sort[i].Val != "" {
				order = order + condiotion.Sort[i].Key + " " + condiotion.Sort[i].Val + ", "
			}
		}

		order = strings.TrimSuffix(order, ", ")

		queryCondition += order
	}

	queryCondition = queryCondition + " LIMIT " + condiotion.Limit
	queryCondition = queryCondition + " OFFSET " + condiotion.Offset

	return queryCondition
}
