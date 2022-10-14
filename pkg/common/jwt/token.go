package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type IToken interface {
	GenerateJwtAccessToken(userId uint) (string, error)
	GenerateJwtRefreshToken(userId uint) (string, error)
}

// https://developer.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
type Token struct{}

type Claims struct {
	UserNo         uint
	StandardClaims jwt.StandardClaims
}

func NewToken() IToken {
	return &Token{}
}

func (t Token) GenerateJwtAccessToken(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t Token) GenerateJwtRefreshToken(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t Token) VerifyJwtTokenValidation() {}
