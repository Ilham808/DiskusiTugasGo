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

func (a *answerRepository) GetUserVote(id uint, idLogin uint) (domain.UserVote, error) {
	var userVote domain.UserVote
	if err := a.db.Where("answer_id = ? AND user_id = ?", id, idLogin).First(&userVote).Error; err != nil {
		return userVote, err
	}

	return userVote, nil
}

func (a *answerRepository) AddUserVote(id uint, idLogin uint, voteType int) error {
	userVote := &domain.UserVote{
		UserID:   idLogin,
		AnswerID: id,
		VoteType: voteType,
	}

	if err := a.db.Create(userVote).Error; err != nil {
		return err
	}

	return nil
}

func (a *answerRepository) UpdateUserVote(id uint, idLogin uint, voteType int) error {
	updatedFields := map[string]interface{}{
		"vote_type": voteType,
	}
	if err := a.db.Model(&domain.UserVote{}).
		Where("answer_id = ? AND user_id = ?", id, idLogin).
		Updates(updatedFields).Error; err != nil {
		return err
	}

	return nil
}
