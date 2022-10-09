package usecase

type IProductUpdateUseCase interface {
	updateProductBySeller() error
}

type productUpdateUseCase struct{}

func NewProductUpdateUseCase() IProductUpdateUseCase {
	return &productUpdateUseCase{}
}

func (uc productUpdateUseCase) updateProductBySeller() error {
	return nil
}
