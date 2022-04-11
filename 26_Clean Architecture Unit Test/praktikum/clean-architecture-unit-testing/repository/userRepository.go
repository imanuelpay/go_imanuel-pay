package repository

import (
	"go-ca/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(DB *gorm.DB) model.UserRepository {
	return &UserRepository{DB}
}

func (u *UserRepository) Create(user *model.User) (*model.User, error) {
	err := u.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAll() (*[]model.User, error) {
	users := make([]model.User, 0)

	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
