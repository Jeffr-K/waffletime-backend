package category

type ICategoryUpdateUseCase interface {
	updateCategoryUseCase() error
}

type categoryUpdateUseCase struct{}

func NewCategoryUpdateUseCase() ICategoryUpdateUseCase {
	return &categoryUpdateUseCase{}
}

func (uc categoryUpdateUseCase) updateCategoryUseCase() error {
	return nil
}
