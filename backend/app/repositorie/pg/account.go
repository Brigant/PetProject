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
func (r AccountDB) InsertAccount(account core.Account) (accountID string, err error) {
	query := `INSERT INTO public.account(
		phone, 
		password, 
		age, 
		role) 
		VALUES ($1, $2, $3, $4) RETURNING id`

	err = r.db.DB.QueryRow(query,
		account.Phone,
		core.SHA256(account.Password, core.Salt),
		account.Age,
		account.Role).Scan(&accountID)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", core.ErrDuplicatePhone
		}

		return "", fmt.Errorf("cannot execute query: %w", err)
	}

	return accountID, nil
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

func (r AccountDB) SelectAccountByID(accountID string) (core.Account, error) {
	var account core.Account

	query := `SELECT id, phone, password, age, role 
	FROM public.account WHERE id=$1`

	err := r.db.DB.QueryRow(query, accountID).Scan(
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
			role, 
			request_host, 
			user_agent, 
			client_ip, 
			expired
		) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING refresh_token`

	err := r.db.DB.QueryRow(query,
		session.AccountID,
		session.Role,
		session.RequestHost,
		session.UserAgent,
		session.ClientIP,
		session.Expired).Scan(&session.RefreshToken)
	if err != nil {
		return core.Session{}, fmt.Errorf("internal error while inserting session: %w", err)
	}

	return session, nil
}

func (r AccountDB) RefreshSession(session core.Session) (core.Session, error) {
	var accointID, role string

	query := `UPDATE public.session 
		SET expired = $1
		WHERE refresh_token=$2 AND request_host=$3 AND user_agent=$4 AND client_ip=$5
		RETURNING account_id, role`

	err := r.db.DB.QueryRow(query, session.Expired,
		session.RefreshToken,
		session.RequestHost,
		session.UserAgent,
		session.ClientIP,
	).Scan(&accointID, &role)
	if err != nil {
		return core.Session{}, fmt.Errorf("can't UPDATE session cuase of: %w", err)
	}

	session.AccountID = accointID
	session.Role = role

	return session, nil
}
