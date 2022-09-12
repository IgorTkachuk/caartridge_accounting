package decom_cause

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/decom_cause"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ decom_cause.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) FindAll(ctx context.Context) ([]decom_cause.DecomCause, error) {
	q := `
		SELECT
			id, name
		FROM
			decommissioning_cause
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			pgError = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("Message: %s; Where: %s; Code: %s, Detail: %s; SqlState: %s", pgError.Message, pgError.Where, pgError.Code, pgError.Detail, pgError.SQLState()))
			return nil, newErr
		}
		return nil, err
	}

	dCauses := make([]decom_cause.DecomCause, 0)
	for rows.Next() {
		var dc decom_cause.DecomCause
		err := rows.Scan(&dc.ID, &dc.Name)
		if err != nil {
			return nil, err
		}
		dCauses = append(dCauses, dc)
	}

	return dCauses, nil
}

func (r repository) FindById(ctx context.Context, id int) (dc decom_cause.DecomCause, err error) {
	q := `
		SELECT
			id, name
		FROM
			decommissioning_Cause
		WHERE
			id=$1
	`
	err = r.client.QueryRow(ctx, q, id).Scan(&dc.ID, &dc.Name)
	if err != nil {
		return decom_cause.DecomCause{}, err
	}

	return
}

func (r repository) Create(ctx context.Context, dto decom_cause.CreateDecomCauseDTO) (id int, err error) {
	q := `
		INSERT INTO
			decommissioning_cause (name)
		VALUES
			($1)
		RETURNING id
	`
	if err = r.client.QueryRow(ctx, q, dto.Name).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf(fmt.Sprintf("Message: %s; Code: %s; Detail: %s; Where: %s; SqlState: %s", pgErr.Message, pgErr.Code, pgErr.Detail, pgErr.Where, pgErr.SQLState()))
			return -1, newErr
		}
		return -1, err
	}

	return
}

func (r repository) Update(ctx context.Context, dto decom_cause.UpdateDecomCauseDTO) error {
	q := `
		UPDATE
			decommissioning_cause
		SET
			name=$1
		WHERE
			id=$2
	`
	_, err := r.client.Exec(ctx, q, &dto.Name, &dto.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
		FROM
			decommissioning_cause
		WHERE
			id=$1
	`

	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) decom_cause.Repository {
	return &repository{client: client, logger: logger}
}
