package pg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nececarry blank import
)

const (
	ErrCodeUniqueViolation     = "unique_violation"
	ErrCodeNoData              = "no_data"
	ErrCodeForeignKeyViolation = "foreign_key_violation"
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

func buildQueryCondition(queryParameter core.ConditionParams) string {
	var queryCondition string

	if len(queryParameter.Filter) > 0 {
		where := "WHERE "

		for i := 0; i < len(queryParameter.Filter); i++ {
			if queryParameter.Filter[i].Val != "" {
				if _, err := strconv.Atoi(queryParameter.Filter[i].Val); err == nil {
					where = where + queryParameter.Filter[i].Key + ">=" + queryParameter.Filter[i].Val + " AND "
				} else {
					where = where + queryParameter.Filter[i].Key + "='" + queryParameter.Filter[i].Val + "' AND "
				}
			}
		}

		where = strings.TrimSuffix(where, "AND ")

		queryCondition += where
	}

	if len(queryParameter.Sort) > 0 {
		order := "ORDER BY "

		for i := 0; i < len(queryParameter.Sort); i++ {
			if queryParameter.Sort[i].Val != "" {
				order = order + queryParameter.Sort[i].Key + " " + queryParameter.Sort[i].Val + ", "
			}
		}

		order = strings.TrimSuffix(order, ", ")

		queryCondition += order
	}

	queryCondition = queryCondition + " LIMIT " + queryParameter.Limit
	queryCondition = queryCondition + " OFFSET " + queryParameter.Offset

	return queryCondition
}
