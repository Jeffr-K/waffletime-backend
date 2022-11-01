package category

type ICategoryDeleteUseCase interface {
	deleteCategoryUsecase() error
}

type categoryDeleteUseCase struct{}

func NewCategoryDeleteUseCase() ICategoryDeleteUseCase {
	return &categoryDeleteUseCase{}
}

func (uc categoryDeleteUseCase) deleteCategoryUsecase() error {
	return nil
}
