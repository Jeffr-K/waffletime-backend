package usecase

import (
	"fmt"
	"waffletime/internal/user/domain"
	"waffletime/internal/user/presentor/dto"
	"waffletime/internal/user/repository"
	"waffletime/pkg/database"
)

type ICreateUseCase interface {
	UserRegisterAsMember(dto *dto.CreateUserDto) error
}

type UserCreateUseCase struct{}

func (uc UserCreateUseCase) UserRegisterAsMember(dto *dto.CreateUserDto) error {
	user := domain.User{Username: dto.Username, Password: dto.Password, Email: dto.Email, Address: dto.Address}
	userRepository := repository.NewRepository(database.Db)
	result := userRepository.SaveTo(&user)
	if result != nil {
		fmt.Println("UserRegisterAsMember >", result)
	}

	return nil
}

//register, err := uc.userAggregateRoot.Register(user)
