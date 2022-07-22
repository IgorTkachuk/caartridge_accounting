package user

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) user.Repository {
	return &repository{
		client: client,
	}
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
