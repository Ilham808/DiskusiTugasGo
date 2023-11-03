package usecase_test

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/domain/mocks"
	"DiskusiTugas/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestSignupUseCase_Store(t *testing.T) {
	mockSignupRepo := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		mockSignupRepo.On("Store", mock.Anything).Return(nil).Once()

		mockData := &domain.User{
			Model: &gorm.Model{
				ID: 1,
			},
			Name:      "test",
			Email:     "test",
			Password:  "test",
			Gender:    "test",
			Avatar:    "test",
			IsStudent: true,
			Status:    "test",
		}
		signupUsecase := usecase.NewSignupUsecase(mockSignupRepo)
		err := signupUsecase.Store(mockData)

		assert.Nil(t, err)
		assert.Equal(t, mockData.ID, uint(1))

		mockSignupRepo.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {
		mockSignupRepo.On("Store", mock.Anything).Return(errors.New("Unexpected")).Once()

		mockData := &domain.User{
			Model: &gorm.Model{
				ID: 1,
			},
			Name:      "test",
			Email:     "test",
			Password:  "test",
			Gender:    "test",
			Avatar:    "test",
			IsStudent: true,
			Status:    "test",
		}
		signupUsecase := usecase.NewSignupUsecase(mockSignupRepo)
		err := signupUsecase.Store(mockData)

		assert.NotNil(t, err)

		mockSignupRepo.AssertExpectations(t)
	})
}

func TestSignupUseCase_GetUserByEmail(t *testing.T) {
	mockSignupRepo := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {

		mockSignupRepo.On("GetByEmail", "test@gmail.com").Return(&domain.User{
			Model: &gorm.Model{
				ID: 1,
			},
			Name:      "test",
			Email:     "test@gmail.com",
			Password:  "test",
			Gender:    "test",
			Avatar:    "test",
			IsStudent: true,
			Status:    "test",
		}, nil).Once()

		signupUsecase := usecase.NewSignupUsecase(mockSignupRepo)
		err := signupUsecase.GetUserByEmail("test@gmail.com")

		assert.NoError(t, err)

		mockSignupRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSignupRepo.On("GetByEmail", "test@gmail.com").Return(&domain.User{}, errors.New("Unexpected")).Once()

		signupUsecase := usecase.NewSignupUsecase(mockSignupRepo)
		err := signupUsecase.GetUserByEmail("test@gmail.com")

		assert.Error(t, err)

		mockSignupRepo.AssertExpectations(t)
	})
}
