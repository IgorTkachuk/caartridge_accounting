package doc_type

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/doc_type"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/jackc/pgconn"
)

var _ doc_type.Repository = &repository{}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) GetAll(ctx context.Context) ([]doc_type.DocType, error) {
	q := `
		SELECT
			id, name
		FROM
			doc_type
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	docTypes := make([]doc_type.DocType, 0)
	for rows.Next() {
		var docType doc_type.DocType
		err := rows.Scan(&docType.ID, &docType.Name)
		if err != nil {
			return nil, err
		}
		docTypes = append(docTypes, docType)
	}

	return docTypes, nil
}

func (r repository) GetById(ctx context.Context, id int) (docType doc_type.DocType, err error) {
	q := `
		SELECT
			id, name
		FROM
			doc_type
		WHERE
			id=$1
	`
	err = r.client.QueryRow(ctx, q, id).Scan(&docType.ID, &docType.Name)
	if err != nil {
		return doc_type.DocType{}, err
	}

	return
}

func (r repository) Create(ctx context.Context, docType doc_type.CreateDocTypeDTO) (id int, err error) {
	q := `
		INSERT INTO doc_type
			(name)
		VALUES
			($1)
		RETURNING id
	`
	if err = r.client.QueryRow(ctx, q, docType.Name).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf(fmt.Sprintf("%s", pgErr.Message))
			return -1, newErr
		}
		return -1, err
	}

	return id, nil
}

func (r repository) Update(ctx context.Context, docType doc_type.UpdateDocTypeDTO) error {
	q := `
		UPDATE
			doc_type
		SET
			name=$1
		WHERE
			id=$2
	`
	_, err := r.client.Exec(ctx, q, docType.Name, docType.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE
		FROM
			doc_type
		WHERE
			id=$1
	`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) doc_type.Repository {
	return &repository{client: client, logger: logger}
}
