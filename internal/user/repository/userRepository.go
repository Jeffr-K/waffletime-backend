package repository

import (
	"errors"
	"gorm.io/gorm"
	"waffletime/internal/user/domain"
)

type UserRepository interface {
	SaveTo(user *domain.User) error
	UpdateTo(user *domain.User) error
	DeleteTo(userId int) error
	GetTo(id int) (domain.User, error)
}

type repository struct {
	database *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &repository{database: db}
}

func (r *repository) SaveTo(user *domain.User) error {
	err := r.database.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateTo(user *domain.User) error {
	err := r.database.Model(&domain.User{}).Where("id = ?", user.ID).Updates(domain.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Address:  user.Address,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTo(userId int) error {
	err := r.database.Delete(&domain.User{}, userId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetTo(id int) (domain.User, error) {
	var result domain.User
	data := r.database.Where("id = ?", uint(id)).First(&result)
	if data.RowsAffected == 0 {
		return domain.User{}, errors.New("user data notFound")
	}
	return result, nil
}

func (r *repository) GetBy(id int, condition string) (domain.User, error) {
	// TODO: implement
	return domain.User{}, nil
}

func (r *repository) GetListBy(condition string) ([]domain.User, error) {
	// TODO: implement
	return []domain.User{}, nil
}

func (r *repository) getAllList() ([]domain.User, error) {
	// TODO: implement
	return []domain.User{}, nil
}
