package ou

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/ou"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ ou.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) FindAll(ctx context.Context) ([]ou.Ou, error) {
	q := `
		SELECT 
			id, name, parent_id, business_line_id
		FROM
			ou
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	ous := make([]ou.Ou, 0)
	for rows.Next() {
		var o ou.Ou
		err = rows.Scan(&o.ID, &o.Name, &o.ParentId, &o.BusinessLineId)
		if err != nil {
			return nil, err
		}
		ous = append(ous, o)
	}

	return ous, nil
}

func (r repository) FindById(ctx context.Context, id int) (o ou.Ou, err error) {
	q := `
		SELECT
			id, name, parent_id, business_line_id
		FROM
			ou
		WHERE
			id=$1
	`

	err = r.client.QueryRow(ctx, q, id).Scan(&o.ID, &o.Name, &o.ParentId, &o.BusinessLineId)
	if err != nil {
		return ou.Ou{}, err
	}

	return
}

func (r repository) Create(ctx context.Context, ou ou.CreateOuDTO) (id int, err error) {
	q := `
		INSERT INTO
			ou
			(name, parent_id, business_line_id)
		VALUES
			($1, $2, $3)
		RETURNING id
	`
	if err = r.client.QueryRow(ctx, q, ou.Name, ou.ParentId, ou.BusinessLineId).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return -1, newErr
		}
		return -1, err
	}

	return
}

func (r repository) Update(ctx context.Context, ou ou.UpdateOuDTO) error {
	q := `
		UPDATE
			ou
		SET
			name=$1, parent_id=$2, business_line_id=$3
		WHERE
			id=$4
	`
	_, err := r.client.Exec(ctx, q, ou.Name, ou.ParentId, ou.BusinessLineId, ou.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
		FROM
			ou
		WHERE
			id=$1
	`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) ou.Repository {
	return &repository{client: client, logger: logger}
}
