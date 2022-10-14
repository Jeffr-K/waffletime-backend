package usecase

import (
	"waffletime/internal/user/domain"
	"waffletime/internal/user/repository"
	"waffletime/pkg/database"
)

type IUserSearchUseCase interface {
	GetUserById(id uint) (domain.User, error)
}

type UserSearchUseCase struct{}

func (us UserSearchUseCase) GetUserById(id int) (domain.User, error) {
	userRepository := repository.NewRepository(database.Db)
	result, err := userRepository.GetTo(id)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (us UserSearchUseCase) GetUserByEmail(email string) (domain.User, error) {
	userRepository := repository.NewRepository(database.Db)
	result, err := userRepository.GetByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}
