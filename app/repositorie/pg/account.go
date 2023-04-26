package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Brigant/PetPorject/app/core"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AccountDB struct {
	db *sqlx.DB
}

func NewAccountDB(db *sqlx.DB) AccountDB {
	return AccountDB{
		db: db,
	}
}

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
		account.Password,
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

func (r AccountDB) SelectSession(session core.Session) (core.Session, error) {
	query := `SELECT  
		refresh_token, account_id, role, request_host, user_agent, client_ip, expired, created
		FROM public.session
		WHERE refresh_token=$1 and request_host=$2 and user_agent=$3 and client_ip=$4`

	err := r.db.DB.QueryRow(
		query,
		session.RefreshToken,
		session.RequestHost,
		session.UserAgent,
		session.ClientIP,
	).Scan(
		&session.RefreshToken,
		&session.AccountID,
		&session.Role,
		&session.RequestHost,
		&session.UserAgent,
		&session.ClientIP,
		&session.Expired,
		&session.Created,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Session{}, core.ErrSesseionNotFound
		}

		return core.Session{}, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return session, nil
}

func (r AccountDB) RefreshSession(session core.Session) error {
	query := `UPDATE public.session 
		SET expired = $1
		WHERE refresh_token=$2
		RETURNING account_id, role`

	_, err := r.db.DB.Exec(query, session.Expired, session.RefreshToken)
	if err != nil {
		return fmt.Errorf("can't UPDATE session cuase of: %w", err)
	}

	return nil
}

func (r AccountDB) DeleteSesions(accountID string) error {
	const minimalRowEffected = 1

	query := `DELETE FROM public.session
		Where account_id=$1`

	result, err := r.db.DB.Exec(query, accountID)
	if err != nil {
		return fmt.Errorf("error while deleting session: %w ", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unexpected error while RowsAffected: %w", err)
	}

	if rowAffected < minimalRowEffected {
		return core.ErrNoRowsEffected
	}

	return nil
}
