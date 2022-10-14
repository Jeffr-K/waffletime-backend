package crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IBCrypt interface {
	HashingPassword(password string) (string, error)
	Match(givenPassword, originPassword string) (bool, error)
}

type bCrypt struct{}

func NewBCrypt() IBCrypt {
	return &bCrypt{}
}

func (b bCrypt) HashingPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		log.Fatalln("hash password is wrong")
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

func (b bCrypt) Match(givenPassword, originPassword string) (bool, error) {
	fmt.Println("givenPassword", givenPassword)
	fmt.Println("originPassword", originPassword)
	err := bcrypt.CompareHashAndPassword([]byte(originPassword), []byte(givenPassword))
	if err != nil {
		log.Fatalln("\n[Password Matching is wrong]\n", err)
		return false, err
	}

	return true, nil
}
