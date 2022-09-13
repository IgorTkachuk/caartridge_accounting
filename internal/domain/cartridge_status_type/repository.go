package cartridge_status_type

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]CartridgeStatusType, error)
	FindById(ctx context.Context, id int) (CartridgeStatusType, error)
	Create(ctx context.Context, dto CreateCartridgeStatusTypeDTO) (int, error)
	Update(ctx context.Context, dto UpdateCartridgeStatusTypeDTO) error
	Delete(ctx context.Context, id int) error
}
