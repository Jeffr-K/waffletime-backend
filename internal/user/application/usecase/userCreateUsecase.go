package usecase

import (
	"fmt"
	"log"
	"waffletime/internal/user/domain"
	"waffletime/internal/user/presentor/dto"
	"waffletime/internal/user/repository"
	BCrypt "waffletime/pkg/common/crypto"
	"waffletime/pkg/database"
)

type ICreateUseCase interface {
	UserRegisterAsMember(dto *dto.CreateUserDto) error
}

type UserCreateUseCase struct{}

func (uc UserCreateUseCase) UserRegisterAsMember(dto *dto.CreateUserDto) error {
	hashedPassword, err := BCrypt.NewBCrypt().HashingPassword(dto.Password)
	if err != nil {
		log.Fatalln("Password Hashing error: ", err)
		return err
	}

	user := domain.User{Username: dto.Username, Password: hashedPassword, Email: dto.Email, Address: dto.Address}
	userRepository := repository.NewRepository(database.Db)
	result := userRepository.SaveTo(&user)
	if result != nil {
		fmt.Println("UserRegisterAsMember >", result)
	}

	return nil
}
