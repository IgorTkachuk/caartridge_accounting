package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/utils"
	"github.com/jackc/pgconn"
)

var _ user.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) FindById(ctx context.Context, id int) (user.User, error) {
	q := `
		SELECT
			id, name, pwf_hash
		FROM
			usr
		WHERE
			id=$1
	`
	r.logger.Trace(fmt.Sprintf("SQL query: %s", utils.FormatQuery(q)))
	var u user.User

	err := r.client.QueryRow(ctx, q, id).Scan(&u.ID, &u.Name, &u.PwdHash)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (r repository) FindByName(ctx context.Context, name string) (user.User, error) {
	q := `
		SELECT
			id, name, pwd_hash
		FROM
			usr
		WHERE
			name=$1
	`
	r.logger.Trace(fmt.Sprintf("SQL query: %s", utils.FormatQuery(q)))
	var u user.User

	err := r.client.QueryRow(ctx, q, name).Scan(&u.ID, &u.Name, &u.PwdHash)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (r repository) FindAll(ctx context.Context) (u []user.User, err error) {
	q := `SELECT
			id, name, pwd_hash
		  FROM
			usr`

	rows, _ := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var usr user.User

		err = rows.Scan(&usr.ID, &usr.Name, &usr.PwdHash)

		users = append(users, usr)
	}

	return users, nil
}

func (r repository) Create(ctx context.Context, u user.User) (int, error) {
	q := `
		INSERT INTO usr
			(name, pwd_hash) 
		VALUES
			($1, $2)
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Qery: %s", utils.FormatQuery(q)))
	if err := r.client.QueryRow(ctx, q, u.Name, u.PwdHash).Scan(&u.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return -1, newErr
		}
		return -1, err
	}

	return u.ID, nil
}
