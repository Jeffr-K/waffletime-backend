package product

type IProductDeleteUseCase interface {
	publishProductBySeller() error
}

type productDeleteUseCase struct{}

func NewProductDeleteUseCase() IProductDeleteUseCase {
	return &productDeleteUseCase{}
}

func (uc productDeleteUseCase) publishProductBySeller() error {
	return nil
}
