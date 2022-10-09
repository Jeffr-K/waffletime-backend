package usecase

import (
	"fmt"
	"github.com/jeffr-k/waffletime/internal/user/repository"
	"github.com/jeffr-k/waffletime/pkg/database"
)

type IUserDeleteUsecase interface {
	UserDropdownMembership(userId int) error
}

type UserDeleteUseCase struct{}

func (ud UserDeleteUseCase) UserDropdownMembership(userId int) error {
	userRepository := repository.NewRepository(database.Db)
	result := userRepository.DeleteTo(userId)
	if result != nil {
		fmt.Println("UserDropdownMembership >", result)
	}

	return nil
}
