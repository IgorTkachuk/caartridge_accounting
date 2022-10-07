package ctr_showcase

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/ctr_showcase"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ ctr_showcase.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) *repository {
	return &repository{client: client, logger: logger}
}

func (r repository) FindAll(ctx context.Context) ([]ctr_showcase.CtrShowcaseDTO, error) {
	q := `
		SELECT 
			id, vendor, model, sn, status, doc_number, 
			doc_date, employee, ou, business_line
		FROM v_ctr	
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	ctrs := make([]ctr_showcase.CtrShowcaseDTO, 0)
	for rows.Next() {
		var ctr ctr_showcase.CtrShowcaseDTO
		err := rows.Scan(
			&ctr.ID, &ctr.Vendor, &ctr.Model, &ctr.Sn, &ctr.Status, &ctr.DocNumber,
			&ctr.DocDate, &ctr.Employee, &ctr.Ou, &ctr.BaseLine,
		)
		if err != nil {
			return nil, err
		}
		ctrs = append(ctrs, ctr)
	}

	return ctrs, nil

}
