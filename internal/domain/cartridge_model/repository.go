package cartridge_model

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]CartridgeModel, error)
	FindById(ctx context.Context, id int) (CartridgeModel, error)
	Create(ctx context.Context, ctrModel CreateCartridgeModelDTO) (int, error)
	Update(ctx context.Context, ctrModel UpdateCartridgeModelDTO) error
	Delete(ctx context.Context, id int) error
}
