package vndr

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/vndr"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/utils"
	"github.com/jackc/pgconn"
)

var _ vndr.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) FindAll(ctx context.Context) (v []vndr.Vendor, err error) {
	q := `
		SELECT
			id, name, logo_url
		FROM
			vendor
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	vendors := make([]vndr.Vendor, 0)
	for rows.Next() {
		var v vndr.Vendor

		rows.Scan(&v.ID, &v.Name, &v.LogoUrl)
		vendors = append(vendors, v)
	}

	return vendors, nil
}

func (r repository) Create(ctx context.Context, v vndr.Vendor) (int, error) {
	q := `
		INSERT INTO vendor
			(name, logo_url)
		VALUES
			($1, $2)
		RETURNING id
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	if err := r.client.QueryRow(ctx, q, v.Name, v.LogoUrl).Scan(&v.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error((newErr))
			return -1, newErr
		}
		return -1, err
	}

	return v.ID, nil
}

func (r repository) FindById(ctx context.Context, id int) (v vndr.Vendor, err error) {
	q := `
		SELECT
			id, name, logo_url
		FROM
			vendor
		WHERE
			id=$1
	`

	if err := r.client.QueryRow(ctx, q, id).Scan(&v.ID, &v.Name, &v.LogoUrl); err != nil {
		return vndr.Vendor{}, err
	}

	return v, nil
}

func (r repository) FindByName(ctx context.Context, name string) ([]vndr.Vendor, error) {
	q := `
		SELECT
			id, name, logo_url
		FROM
			vendor
		WHERE 
			name LIKE '%$1%'
	`
	rows, err := r.client.Query(ctx, q, name)
	if err != nil {
		return nil, err
	}

	vendors := make([]vndr.Vendor, 0)

	for rows.Next() {
		var v vndr.Vendor

		rows.Scan(&v.ID, &v.Name, &v.LogoUrl)
		vendors = append(vendors, v)
	}

	return vendors, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) vndr.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
