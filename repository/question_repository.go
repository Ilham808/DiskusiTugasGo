package repository

import (
	"DiskusiTugas/domain"

	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domain.QuestionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) FetchWithPagination(page, pageSize int) ([]domain.Question, int, error) {
	var question []domain.Question
	var totalItems int64

	if err := q.db.Model(&domain.Question{}).
		Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := q.db.
		Preload("User").
		Preload("Subject").
		Select("id, created_at, updated_at, deleted_at, user_id, subject_id, question, description, file").
		Offset(offset).
		Limit(pageSize).
		Find(&question).Error; err != nil {
		return nil, 0, err
	}

	return question, int(totalItems), nil
}

func (q *questionRepository) GetByID(id int) (domain.Question, error) {
	var question domain.Question

	if err := q.db.Preload("User").Preload("Answer.User").Preload("Answer.Comment.User").Where("id = ?", id).First(&question).Error; err != nil {
		return question, err
	}

	return question, nil

}

func (q *questionRepository) Store(question *domain.Question) error {
	if err := q.db.Create(question).Error; err != nil {
		return err
	}
	return nil
}
