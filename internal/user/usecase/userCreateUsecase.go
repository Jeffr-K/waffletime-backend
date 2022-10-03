package usecase

import (
	"github.com/jeffr-k/waffletime/internal/user/domain"
	"github.com/jeffr-k/waffletime/internal/user/presentor/dto"
)

type UserCreateUseCase struct{}

func (uc UserCreateUseCase) UserRegisterAsMember(dto *dto.CreateUserDto) (domain.User, error) {
	return domain.User{}, nil
}
