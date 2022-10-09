package usecase

import (
	"github.com/jeffr-k/waffletime/internal/user/domain"
	"github.com/jeffr-k/waffletime/internal/user/repository"
	"github.com/jeffr-k/waffletime/pkg/database"
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
