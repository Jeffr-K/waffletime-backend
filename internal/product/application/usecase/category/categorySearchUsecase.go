package category

type ICategorySearchUseCase interface {
	searchCategoryUsecase() error
}

type categorySearchUsecase struct{}

func NewCategorySearchUseCase() ICategorySearchUseCase {
	return &categorySearchUsecase{}
}

func (uc categorySearchUsecase) searchCategoryUsecase() error {
	return nil
}
