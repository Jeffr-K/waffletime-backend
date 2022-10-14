package usecase

import (
	"log"
	req "waffletime/internal/auth/presentor/dto/req"
	res "waffletime/internal/auth/presentor/dto/res"
	"waffletime/internal/user/application/usecase"
	BCrypt "waffletime/pkg/common/crypto"
	"waffletime/pkg/common/jwt"
)

type IUserLoginUsecase interface {
	Login(dto *req.UserLoginDto) (res.UserLoginResponseDto, error)
	Logout()
}

type UserLoginUsecase struct {
	userSearchUseCase usecase.UserSearchUseCase
}

func (ul UserLoginUsecase) Login(dto *req.UserLoginDto) (res.UserLoginResponseDto, error) {

	user, err := ul.userSearchUseCase.GetUserByEmail(dto.Email)
	if err != nil {
		log.Fatalln("User Not Found in Login Service", err)
		return res.UserLoginResponseDto{
			StatusCode: 404,
			Message:    "User Not Found",
			Tokens: res.TokensMap{
				AccessToken:  "",
				RefreshToken: "",
			}}, err
	}

	Match, err := BCrypt.NewBCrypt().Match(dto.Password, user.Password)
	if err != nil {
		log.Fatalln("<Login>Error/PasswordMatchError: ", err)
		return res.UserLoginResponseDto{
			StatusCode: 404,
			Message:    "User Not Found",
			Tokens: res.TokensMap{
				AccessToken:  "",
				RefreshToken: "",
			}}, err
	}
	if Match == false {
		log.Fatalln("\n<Login>Error/PasswordNotMatched: \n", err)
		return res.UserLoginResponseDto{
			StatusCode: 404,
			Message:    "User Not Found",
			Tokens: res.TokensMap{
				AccessToken:  "",
				RefreshToken: "",
			}}, err
	}

	token := jwt.NewToken()
	accessToken, err := token.GenerateJwtAccessToken(user.ID)
	if err != nil {
		log.Fatalln("\n[<Login>GenerateRefreshToken]\n", err)
		return res.UserLoginResponseDto{
			StatusCode: 404,
			Message:    "User Not Found",
			Tokens: res.TokensMap{
				AccessToken:  "",
				RefreshToken: "",
			}}, err
	}
	refreshToken, err := token.GenerateJwtRefreshToken(user.ID)
	if err != nil {
		log.Fatalln("\n[<Login>GenerateRefreshToken]\n", err)
		return res.UserLoginResponseDto{
			StatusCode: 404,
			Message:    "User Not Found",
			Tokens: res.TokensMap{
				AccessToken:  "",
				RefreshToken: "",
			}}, err
	}

	return res.UserLoginResponseDto{
		StatusCode: 204,
		Message: "User Successfully login",
		Tokens: res.TokensMap{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}}, err
}

func (ul UserLoginUsecase) Logout() {}
