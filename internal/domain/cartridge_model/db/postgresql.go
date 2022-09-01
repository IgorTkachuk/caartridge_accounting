package cartridge_model

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_model"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ cartridge_model.Repository = &repository{}

type repository struct {
	logger *logging.Logger
	client postgresql.Client
}

func NewRepository(client postgresql.Client, logger *logging.Logger) cartridge_model.Repository {
	return &repository{
		logger: logger,
		client: client,
	}
}

func (r repository) FindAll(ctx context.Context) ([]cartridge_model.CartridgeModel, error) {
	q := `
		SELECT
			id, name, vendor_id, image_url
		FROM
			ctr_model
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	ctrModels := make([]cartridge_model.CartridgeModel, 0)

	for rows.Next() {
		var ctr cartridge_model.CartridgeModel
		rows.Scan(&ctr.ID, &ctr.Name, &ctr.VendorId, &ctr.ImageUrl)
		qSuppPrns := `
		SELECT
			prn_model_id
		FROM
			ctr_supp_prn
		WHERE
			ctr_model_id=$1
		`
		rows, err := r.client.Query(ctx, qSuppPrns, ctr.ID)
		if err != nil {
			return nil, err
		}

		suppPrns := make([]int, 0)
		for rows.Next() {
			var suppPrnId int
			rows.Scan(&suppPrnId)
			suppPrns = append(suppPrns, suppPrnId)
		}
		ctr.SuppPrns = suppPrns

		ctrModels = append(ctrModels, ctr)
	}

	return ctrModels, nil
}

func (r repository) FindById(ctx context.Context, id int) (c cartridge_model.CartridgeModel, err error) {
	q := `
		SELECT
			id, name, vendor_id, image_url
		FROM 
			ctr_model
		WHERE
			id=$1
	`

	err = r.client.QueryRow(ctx, q, id).Scan(&c.ID, &c.Name, &c.VendorId, &c.ImageUrl)
	if err != nil {
		return cartridge_model.CartridgeModel{}, err
	}

	qSuppPrns := `
		SELECT 
			prn_model_id
		FROM
			ctr_supp_prn
		WHERE
			ctr_model_id=$1
	`

	rows, err := r.client.Query(ctx, qSuppPrns, id)
	if err != nil {
		return cartridge_model.CartridgeModel{}, err
	}

	suppPrnModels := make([]int, 0)
	for rows.Next() {
		var suppPrnModel int
		rows.Scan(&suppPrnModel)
		suppPrnModels = append(suppPrnModels, suppPrnModel)
	}
	c.SuppPrns = suppPrnModels

	return
}

func (r repository) Create(ctx context.Context, ctrModel cartridge_model.CreateCartridgeModelDTO) (id int, err error) {
	q := `
		INSERT INTO 
			ctr_model(name, vendor_id, image_url)
		VALUES
			($1, $2, $3)
		RETURNING id
	`

	if err = r.client.QueryRow(ctx, q, &ctrModel.Name, &ctrModel.VendorId, &ctrModel.ImageUrl).Scan(&id); err != nil {
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

func (r repository) Update(ctx context.Context, ctrModel cartridge_model.UpdateCartridgeModelDTO) error {
	q := `
		UPDATE
			ctr_model
		SET
			name=$1, vendor_id=$2, image_url=$3
		WHERE
			id=$4
	`

	ct, err := r.client.Exec(ctx, q, &ctrModel.Name, &ctrModel.VendorId, &ctrModel.ImageUrl, &ctrModel.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Carteidge model update operation affected rows: %d", ct.RowsAffected())

	qSuppPrnsDel := `
		DELETE
		FROM 
			ctr_supp_prn
		WHERE
			ctr_model_id=$1
	`
	ct, err = r.client.Exec(ctx, qSuppPrnsDel, &ctrModel.ID)
	if err != nil {
		return err
	}

	if len(ctrModel.SuppPrns) == 0 {
		return nil
	}

	qSuppPrnsIns := `
		INSERT INTO
			ctr_supp_prn (ctr_model_id, prn_model_id)
		VALUES

	`
	for i, prn := range ctrModel.SuppPrns {
		row := fmt.Sprintf(" (%d,%d)", ctrModel.ID, prn)
		qSuppPrnsIns += row
		if i != len(ctrModel.SuppPrns)-1 {
			qSuppPrnsIns += ", "
		}
	}

	ct, err = r.client.Exec(ctx, qSuppPrnsIns)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
		FROM 
			ctr_model
		WHERE
			id=$1
	`

	ct, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	fmt.Printf("CommandTag (RowAffected) ####: %d", ct.RowsAffected())
	return nil
}
