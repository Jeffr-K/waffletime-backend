package product

type IProductPublishUseCase interface {
	publishProductBySeller() error
}

type productPublishUseCase struct{}

func NewProductPublishUseCase() IProductPublishUseCase {
	return &productPublishUseCase{}
}

func (uc productPublishUseCase) publishProductBySeller() error {
	return nil
}
