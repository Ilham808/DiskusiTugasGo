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

func TestSubjectUseCase_Fetch(t *testing.T) {
	mockSubjectRepo := new(mocks.SubjectRepository)

	t.Run("success", func(t *testing.T) {

		mockSubjectData := []domain.Subject{
			{
				Model: &gorm.Model{
					ID: uint(1),
				},
				Name: "test",
				Slug: "test",
			},
			{
				Model: &gorm.Model{
					ID: uint(2),
				},
				Name: "test",
				Slug: "test",
			},
		}
		mockSubjectRepo.On("Fetch").Return(mockSubjectData, nil).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		subjects, err := subjectUseCase.Fetch()

		assert.NoError(t, err)
		assert.NotNil(t, subjects)

		mockSubjectRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSubjectRepo.On("Fetch").Return(nil, errors.New("Unexpected")).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		subjects, err := subjectUseCase.Fetch()

		assert.Error(t, err)
		assert.Nil(t, subjects)

		mockSubjectRepo.AssertExpectations(t)
	})
}

func TestSubjectUseCase_GetByID(t *testing.T) {
	mockSubjectRepo := new(mocks.SubjectRepository)

	t.Run("success", func(t *testing.T) {
		mockSubjectRepo.On("GetByID", mock.Anything).Return(domain.Subject{
			Model: &gorm.Model{
				ID: uint(1),
			},
			Name: "test",
			Slug: "test",
		}, nil).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		subject, err := subjectUseCase.GetByID(1)

		assert.NoError(t, err)
		assert.NotNil(t, subject)

		mockSubjectRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSubjectRepo.On("GetByID", mock.Anything).Return(domain.Subject{}, errors.New("Unexpected")).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		subject, err := subjectUseCase.GetByID(1)
		expectData := domain.Subject{}

		assert.Error(t, err)
		assert.Equal(t, expectData, subject)

		mockSubjectRepo.AssertExpectations(t)
	})
}

func TestSubjectUseCase_Store(t *testing.T) {
	mockSubjectRepo := new(mocks.SubjectRepository)

	t.Run("success", func(t *testing.T) {
		mockSubjectRepo.On("Store", mock.Anything).Return(nil).Once()

		mockData := &domain.Subject{
			Name: "test",
			Slug: "test",
		}
		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Store(mockData)

		assert.NoError(t, err)
		assert.Equal(t, mockData, mockData)

		mockSubjectRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSubjectRepo.On("Store", mock.Anything).Return(errors.New("Unexpected")).Once()

		mockData := &domain.Subject{
			Name: "test",
			Slug: "test",
		}

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Store(mockData)

		assert.Error(t, err)

		mockSubjectRepo.AssertExpectations(t)
	})
}

func TestSubjectUseCase_Update(t *testing.T) {
	mockSubjectRepo := new(mocks.SubjectRepository)

	t.Run("success", func(t *testing.T) {
		mockSubjectRepo.On("Update", mock.Anything).Return(nil).Once()

		mockData := &domain.Subject{
			Name: "test",
			Slug: "test",
		}

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Update(mockData)

		assert.NoError(t, err)

		mockSubjectRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSubjectRepo.On("Update", mock.Anything).Return(errors.New("Unexpected")).Once()

		mockData := &domain.Subject{
			Name: "test",
			Slug: "test",
		}

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Update(mockData)

		assert.Error(t, err)

		mockSubjectRepo.AssertExpectations(t)
	})
}

func TestSubjectUseCase_Destroy(t *testing.T) {
	mockSubjectRepo := new(mocks.SubjectRepository)

	t.Run("success", func(t *testing.T) {
		mockSubjectRepo.On("Destroy", mock.Anything).Return(nil).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Destroy(1)

		assert.NoError(t, err)

		mockSubjectRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSubjectRepo.On("Destroy", mock.Anything).Return(errors.New("Unexpected")).Once()

		subjectUseCase := usecase.NewSubjectUseCase(mockSubjectRepo)
		err := subjectUseCase.Destroy(1)

		assert.Error(t, err)

		mockSubjectRepo.AssertExpectations(t)
	})

}
