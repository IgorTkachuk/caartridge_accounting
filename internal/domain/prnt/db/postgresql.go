package prnt

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/prnt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ prnt.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) prnt.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) FindAll(ctx context.Context) ([]prnt.Prn, error) {
	q := `
		SELECT 
			id, name, vendor_id, image_url
		FROM
			prn_model
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	printers := make([]prnt.Prn, 0)

	for rows.Next() {
		var p prnt.Prn
		rows.Scan(&p.ID, &p.Name, &p.VendorID, &p.ImageUrl)
		printers = append(printers, p)
	}

	return printers, nil
}

func (r repository) Create(ctx context.Context, prn prnt.CreatePrnDTO) (id int, err error) {
	q := `
		INSERT INTO 
			prn_model (name, vendor_id, image_url)
		VALUES 
			($1, $2, $3)
		RETURNING id
	`

	if err := r.client.QueryRow(ctx, q, prn.Name, prn.VendorID, prn.ImageUrl).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return -1, newErr
		}
		return -1, err
	}

	return
}

func (r repository) FindById(ctx context.Context, id int) (prn prnt.Prn, err error) {
	q := `
		SELECT
			id, name, vendor_id, image_url
		FROM
			prn_model
		WHERE
			id=$1
	`

	if err = r.client.QueryRow(ctx, q, id).Scan(&prn); err != nil {
		return prnt.Prn{}, err
	}

	return
}

func (r repository) FindByName(ctx context.Context, name string) ([]prnt.Prn, error) {
	q := `
		SELECT
			id, name, vendor_id, image_url
		FROM
			prn_model
		WHERE
			name LIKE '%$1%'
	`

	rows, err := r.client.Query(ctx, q, name)
	if err != nil {
		return nil, err
	}

	printers := make([]prnt.Prn, 0)
	for rows.Next() {
		var printer prnt.Prn
		rows.Scan(&printer.ID, &printer.Name, &printer.VendorID, &printer.ImageUrl)
		printers = append(printers, printer)
	}

	return printers, nil
}

func (r repository) Delete(ctx context.Context, id int) (int, error) {
	q := `
		DELETE
		FROM
			prn_model
		WHERE
			id=$1
	`

	ct, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return 0, err
	}
	fmt.Printf("CommandTag (RowAffected) ####: %d", ct.RowsAffected())

	return int(ct.RowsAffected()), nil
}

func (r repository) Update(ctx context.Context, prn prnt.UpdatePrnDTO) error {
	q := `
		UPDATE
			prn_model
		SET
			name=$1, vendor_id=$2, image_url=$3
		WHERE
			id=$4
	`
	ct, err := r.client.Exec(ctx, q, &prn.Name, &prn.VendorID, &prn.ImageUrl, &prn.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Vendor update delete operation affected rows: %d", ct.RowsAffected())
	return nil
}
