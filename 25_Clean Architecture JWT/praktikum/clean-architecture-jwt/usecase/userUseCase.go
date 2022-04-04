package usecase

import "go-ca/model"

type UserUseCase struct {
	userRepository model.UserRepository
}

func CreateUserUsecase(userRepository model.UserRepository) model.UserUseCase {
	return &UserUseCase{userRepository}
}

func (u *UserUseCase) Create(user *model.User) (*model.User, error) {
	return u.userRepository.Create(user)
}

func (u *UserUseCase) GetByEmailAndPassword(request *model.UserLoginRequest) (*model.User, error) {
	return u.userRepository.GetByEmailAndPassword(request)
}

func (u *UserUseCase) GetAll() (*[]model.User, error) {
	return u.userRepository.GetAll()
}
