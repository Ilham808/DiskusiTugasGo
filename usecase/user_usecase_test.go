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

func TestUserUsecase_Fetch(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	UserID := 1

	t.Run("success", func(t *testing.T) {

		mockUserRepository.On("Fetch").Return([]domain.User{
			{
				Model: &gorm.Model{
					ID: uint(UserID),
				},
				Name:      "test",
				Email:     "test",
				Password:  "test",
				Gender:    "test",
				Avatar:    "test",
				IsStudent: true,
				Status:    "test",
			},
		}, nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, err := userUsecase.Fetch()

		assert.NoError(t, err)
		assert.NotNil(t, users)

		mockUserRepository.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, err := userUsecase.Fetch()

		assert.Error(t, err)
		assert.Nil(t, users)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_FetchWithPagination(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {

		mockUserData := []domain.User{
			{
				Model: &gorm.Model{
					ID: uint(1),
				},
				Name:      "test",
				Email:     "test",
				Password:  "test",
				Gender:    "test",
				Avatar:    "test",
				IsStudent: true,
				Status:    "test",
			},
			{
				Model: &gorm.Model{
					ID: uint(2),
				},
				Name:      "test",
				Email:     "test",
				Password:  "test",
				Gender:    "test",
				Avatar:    "test",
				IsStudent: true,
				Status:    "test",
			},
		}

		mockUserRepository.On("FetchWithPagination", mock.Anything, mock.Anything).Return(mockUserData, 2, nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, totalItems, err := userUsecase.FetchWithPagination(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, len(mockUserData), totalItems)
		assert.NotNil(t, users)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("FetchWithPagination", mock.Anything, mock.Anything).Return(nil, 0, errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, totalItems, err := userUsecase.FetchWithPagination(1, 1)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, totalItems)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_GetByEmail(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {

		expectedUser := &domain.User{
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
		}

		mockUserRepository.On("GetByEmail", "test@gmail.com").Return(expectedUser, nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		user := userUsecase.GetByEmail("test@gmail.com")

		assert.NoError(t, user)
		assert.Equal(t, "test@gmail.com", expectedUser.Email)

		mockUserRepository.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("GetByEmail", "test@gmail.com").Return(nil, errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		user := userUsecase.GetByEmail("test@gmail.com")

		assert.Error(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_GetByID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	UserID := 1

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.Anything).Return(domain.User{
			Model: &gorm.Model{
				ID: uint(UserID),
			},
			Name:      "test",
			Email:     "test",
			Password:  "test",
			Gender:    "test",
			Avatar:    "test",
			IsStudent: true,
			Status:    "test",
		}, nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		user, err := userUsecase.GetByID(UserID)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("GetByID", UserID).Return(domain.User{}, errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		user, err := userUsecase.GetByID(UserID)
		expectedUser := domain.User{}

		assert.Error(t, err)
		assert.Equal(t, expectedUser, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_Update(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything).Return(nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.Update(&domain.User{
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
		})

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything).Return(errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.Update(&domain.User{
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
		})

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_Delete(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Destroy", mock.Anything).Return(nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.Destroy(1)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("Destroy", mock.Anything).Return(errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.Destroy(1)

		assert.Error(t, err)
	})
}

func TestUserUsecase_FecthStudent(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {

		mockUserData := []domain.User{
			{
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
			},
			{
				Model: &gorm.Model{
					ID: uint(2),
				},
				Name:      "test",
				Email:     "test",
				Password:  "test",
				Gender:    "test",
				Avatar:    "test",
				IsStudent: true,
				Status:    "test",
			},
		}
		mockUserRepository.On("FecthStudent", mock.Anything, mock.Anything).Return(mockUserData, 2, nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, totalItems, err := userUsecase.FecthStudent(1, 1)

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 2, totalItems)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("FecthStudent", mock.Anything, mock.Anything).Return([]domain.User{}, 0, errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		users, totalItems, err := userUsecase.FecthStudent(1, 1)
		// expectedUser := []domain.User{}

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, totalItems)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserUsecase_BlockStudent(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("BlockStudent", mock.Anything).Return(nil).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.BlockStudent(1)

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("BlockStudent", mock.Anything).Return(errors.New("Unexpected")).Once()

		userUsecase := usecase.NewUserUseCase(mockUserRepository)
		err := userUsecase.BlockStudent(1)

		assert.Error(t, err)

		mockUserRepository.AssertExpectations(t)
	})
}
