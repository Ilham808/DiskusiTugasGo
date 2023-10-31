package repository

import (
	"DiskusiTugas/domain"

	"gorm.io/gorm"
)

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) domain.AnswerRepository {
	return &answerRepository{
		db: db,
	}
}

func (a *answerRepository) Store(answer domain.Answer) (domain.Answer, error) {
	if err := a.db.Create(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}

func (a *answerRepository) GetByID(id int) (domain.Answer, error) {
	var answer domain.Answer
	if err := a.db.First(&answer, id).Error; err != nil {
		return answer, err
	}

	return answer, nil
}

func (a *answerRepository) Update(id int, answer *domain.Answer) error {
	if err := a.db.Model(&answer).Where("id = ?", id).Updates(&answer).Error; err != nil {
		return err
	}
	return nil
}

func (a *answerRepository) Destroy(id int) error {
	if err := a.db.Where("id = ?", id).Delete(&domain.Answer{}).Error; err != nil {
		return err
	}

	return nil
}

func (a *answerRepository) GetQuestionByID(id uint) (domain.Question, error) {
	var question domain.Question
	if err := a.db.First(&question, id).Error; err != nil {
		return question, err
	}

	return question, nil
}
