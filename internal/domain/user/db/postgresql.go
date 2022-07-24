package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
	"strings"
)

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

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", "")
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

func (r repository) Create(ctx context.Context, u user.User) error {
	q := `
		INSERT INTO usr
			(name, pwd_hash) 
		VALUES
			($1, $2)
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Qery: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, u.Name, u.PwdHash).Scan(&u.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQ: Error: %s, Detail: %s, Where: %s, Cde: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}
