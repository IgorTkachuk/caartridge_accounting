package doc_type

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]DocType, error)
	GetById(ctx context.Context, id int) (DocType, error)
	Create(ctx context.Context, docType CreateDocTypeDTO) (int, error)
	Update(ctx context.Context, docType UpdateDocTypeDTO) error
	Delete(ctx context.Context, id int) error
}
