package usecase_test

import (
	"errors"
	"go-ca/model"
	"go-ca/model/mocks"
	"go-ca/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUserData := model.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "test",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Create", mock.Anything).Return(&mockUserData, nil).Once()

		userUseCase := usecase.CreateUserUsecase(mockUserRepository)
		user, err := userUseCase.Create(&mockUserData)

		assert.NoError(t, err)
		assert.Equal(t, mockUserData, *user)
	})

	t.Run("failed", func(t *testing.T) {
		mockUserRepository.On("Create", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		userUseCase := usecase.CreateUserUsecase(mockUserRepository)
		_, err := userUseCase.Create(&mockUserData)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	users := new([]model.User)
	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("GetAll", mock.Anything).Return(users, nil).Once()

		userUseCase := usecase.CreateUserUsecase(mockUserRepository)
		_, err := userUseCase.GetAll()

		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockUserRepository.On("GetAll", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		userUseCase := usecase.CreateUserUsecase(mockUserRepository)
		_, err := userUseCase.GetAll()

		assert.Error(t, err)
	})
}
