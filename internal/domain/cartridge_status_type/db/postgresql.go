package cartridge_status_type

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_status_type"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ cartridge_status_type.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) FindAll(ctx context.Context) ([]cartridge_status_type.CartridgeStatusType, error) {
	q := `
		SELECT
			id, name
		FROM ctr_status_type
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("Message: %s; Code: %s; Where: %s; Detail: %s, SqlStatus: %s", pgErr.Message, pgErr.Code, pgErr.Where, pgErr.Detail, pgErr.SQLState()))
			return nil, newErr
		}
		return nil, err
	}

	cStatuses := make([]cartridge_status_type.CartridgeStatusType, 0)
	for rows.Next() {
		var cStatus cartridge_status_type.CartridgeStatusType
		err := rows.Scan(&cStatus.ID, &cStatus.Name)
		if err != nil {
			return nil, err
		}
		cStatuses = append(cStatuses, cStatus)
	}

	return cStatuses, nil
}

func (r repository) FindById(ctx context.Context, id int) (cStatus cartridge_status_type.CartridgeStatusType, err error) {
	q := `
		SELECT
			id, name
		FROM
			ctr_status_type
		WHERE
			id=$1
	`

	err = r.client.QueryRow(ctx, q, id).Scan(&cStatus.ID, &cStatus.Name)
	if err != nil {
		return cartridge_status_type.CartridgeStatusType{}, err
	}

	return
}

func (r repository) Create(ctx context.Context, dto cartridge_status_type.CreateCartridgeStatusTypeDTO) (id int, err error) {
	q := `
		INSERT INTO
			ctr_status_type
				(name)
		VALUES
			($1)
		RETURNING id
	`
	err = r.client.QueryRow(ctx, q, dto.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return
}

func (r repository) Update(ctx context.Context, dto cartridge_status_type.UpdateCartridgeStatusTypeDTO) error {
	q := `
		UPDATE
			ctr_status_type
		SET 
			name=$1
		WHERE
			id=$2
	`
	_, err := r.client.Exec(ctx, q, dto.Name, dto.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
		FROM
			ctr_status_type
		WHERE
			id=$1
	`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) *repository {
	return &repository{client: client, logger: logger}
}
