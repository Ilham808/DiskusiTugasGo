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
	if err := q.db.Offset(offset).
		Limit(pageSize).
		Find(&question).Error; err != nil {
		return nil, 0, err
	}

	return question, int(totalItems), nil
}

func (q *questionRepository) Store(question *domain.Question) error {
	if err := q.db.Create(question).Error; err != nil {
		return err
	}
	return nil
}
