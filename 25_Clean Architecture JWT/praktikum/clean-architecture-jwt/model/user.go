package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserUseCase interface {
	Create(user *User) (*User, error)
	GetByEmailAndPassword(request *UserLoginRequest) (*User, error)
	GetAll() (*[]User, error)
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByEmailAndPassword(request *UserLoginRequest) (*User, error)
	GetAll() (*[]User, error)
}
