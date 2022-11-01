package category

type ICategoryCreateUseCase interface {
	createCategoryUsecase() error
}

type categoryCreateUseCase struct{}

func NewProductPublishUseCase() ICategoryCreateUseCase {
	return &categoryCreateUseCase{}
}

func (uc categoryCreateUseCase) createCategoryUsecase() error {
	return nil
}
