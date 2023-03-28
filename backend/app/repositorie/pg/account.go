package pg

import (
	"database/sql"
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

	query := `INSERT INTO public.account(
		phone, 
		password, 
		age, 
		role) 
		VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.DB.QueryRow(query,
		account.Phone,
		core.SHA256(account.Password, core.Salt),
		account.Age,
		account.Role).Scan(&id)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", core.ErrDuplicatePhone
		}

		return "", fmt.Errorf("cannot execute query: %w", err)
	}

	return id, nil
}

func (r AccountDB) SelectAccountByPhone(phone string) (core.Account, error) {
	var account core.Account

	query := `SELECT id, phone, password, age, role 
		FROM public.account WHERE phone=$1`

	err := r.db.DB.QueryRow(query, phone).Scan(
		&account.ID,
		&account.Phone,
		&account.Password,
		&account.Age,
		&account.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Account{}, core.ErrUserNotFound
		}

		return core.Account{}, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return account, nil
}

func (r AccountDB) InsertSession(session core.Session) (core.Session, error) {
	query := `INSERT INTO public.session(
			account_id, 
			request_host, 
			user_agent, 
			client_ip, 
			is_blocked, 
			expired_in
		) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING refresh_token`

	err := r.db.DB.QueryRow(query,
		session.AccountID,
		session.RequestHost,
		session.UserAgent,
		session.ClientIP,
		false,
		session.ExpiredIn).Scan(&session.RefreshToken)
	if err != nil {
		return core.Session{}, fmt.Errorf("internal error while inserting session: %w", err)
	}

	return session, nil
}
