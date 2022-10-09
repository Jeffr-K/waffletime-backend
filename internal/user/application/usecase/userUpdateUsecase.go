package usecase

import (
	"github.com/jeffr-k/waffletime/internal/user/domain"
	"github.com/jeffr-k/waffletime/internal/user/presentor/dto"
	"github.com/jeffr-k/waffletime/internal/user/repository"
	"github.com/jeffr-k/waffletime/pkg/database"
)

type IUserUpdateUseCase interface {
	UpdateUserInfo(user *dto.UpdateUserDto) error
}

type UserUpdateUseCase struct{}

func (uu UserUpdateUseCase) UpdateUserInfo(dto *dto.UpdateUserDto) error {
	//cast, _ := strconv.Atoi(dto.ID)
	user := domain.User{ID: dto.ID, Username: dto.Username, Password: dto.Password, Email: dto.Email, Address: dto.Address}
	userRepository := repository.NewRepository(database.Db)
	err := userRepository.UpdateTo(&user)
	if err != nil {
		return err
	}

	return nil
}
