package business_line

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/business_line"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ business_line.Repository = &repository{}

type repository struct {
	logger *logging.Logger
	client postgresql.Client
}

func (r repository) FindAll(ctx context.Context) ([]business_line.BusinessLine, error) {
	q := `
		SELECT 
			id, name
		FROM 
			business_line	
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	bls := make([]business_line.BusinessLine, 0)
	for rows.Next() {
		var bl business_line.BusinessLine
		err := rows.Scan(&bl.ID, &bl.Name)
		if err != nil {
			return nil, err
		}
		bls = append(bls, bl)
	}

	return bls, nil
}

func (r repository) FindById(ctx context.Context, id int) (bl business_line.BusinessLine, err error) {
	q := `
		SELECT
			id, name
		FROM
			business_line
		WHERE
			id=$1
	`
	err = r.client.QueryRow(ctx, q, id).Scan(&bl.ID, &bl.Name)
	if err != nil {
		return business_line.BusinessLine{}, err
	}

	return bl, nil
}

func (r repository) Create(ctx context.Context, bl business_line.CreateBusinessLineDTO) (id int, err error) {
	q := `
		INSERT INTO
			business_line
				(name)
		VALUES
			($1)
		RETURNING id
	`

	if err = r.client.QueryRow(ctx, q, &bl.Name).Scan(&id); err != nil {
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

func (r repository) Update(ctx context.Context, bl business_line.UpdateBusinessLineDTO) error {
	q := `
		UPDATE
			business_line
		SET
			name=$1
		WHERE
			id=$2
	`

	_, err := r.client.Exec(ctx, q, bl.Name, bl.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
			FROM
				business_line
		WHERE
			id=$1
	`

	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) business_line.Repository {
	return &repository{logger: logger, client: client}
}
