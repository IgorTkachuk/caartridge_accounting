package employee

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/employee"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ employee.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) employee.Repository {
	return &repository{client: client, logger: logger}
}

func (r repository) FindAll(ctx context.Context) ([]employee.Employee, error) {
	q := `
		SELECT 
			id, name, ou_id
		FROM
			employee
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	es := make([]employee.Employee, 0)
	for rows.Next() {
		var e employee.Employee
		err := rows.Scan(&e.ID, &e.Name, &e.OuID)
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}

	return es, nil
}

func (r repository) FindById(ctx context.Context, id int) (e employee.Employee, err error) {
	q := `
		SELECT
			id, name, ou_id
		FROM
			employee
		WHERE
			id=$1
	`
	err = r.client.QueryRow(ctx, q, id).Scan(&e.ID, &e.Name, &e.OuID)
	if err != nil {
		return employee.Employee{}, err
	}

	return
}

func (r repository) Create(ctx context.Context, e employee.CreateEmployeeDTO) (id int, err error) {
	q := `
		INSERT INTO
			employee
				(name, ou_id)
		VALUES
			($1, $2)
		RETURNING id
	`

	if err = r.client.QueryRow(ctx, q, e.Name, e.OuId).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("%s", pgErr.Message))
			return -1, newErr
		}
		return -1, err
	}

	return
}

func (r repository) Update(ctx context.Context, e employee.UpdateEmployeeDTO) error {
	q := `
		UPDATE employee
		SET 
			name=$1, ou_id=$2
		WHERE
			id=$3
	`
	_, err := r.client.Exec(ctx, q, e.Name, e.OuId, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM 
			employee
		WHERE
			id=$1
	`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
