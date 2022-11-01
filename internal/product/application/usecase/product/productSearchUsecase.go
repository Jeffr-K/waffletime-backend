package product

import "waffletime/internal/product/domain"

type IProductSearchUseCase interface {
	GetOneProduct() (domain.Product, error)
	GetALlProductList() (domain.Product, error)
	GetAllProductBy() ([]domain.Product, error)
}

type productSearchUseCase struct{}

func NewProductSearchUseCase() IProductSearchUseCase {
	return &productSearchUseCase{}
}

func (p productSearchUseCase) GetOneProduct() (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productSearchUseCase) GetALlProductList() (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productSearchUseCase) GetAllProductBy() ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
