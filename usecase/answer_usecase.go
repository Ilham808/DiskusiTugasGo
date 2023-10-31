package usecase

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
	"errors"
)

type answerUseCase struct {
	answerRepository domain.AnswerRepository
}

func NewAnswerUseCase(answerRepository domain.AnswerRepository) domain.AnswerUsecase {
	return &answerUseCase{
		answerRepository: answerRepository,
	}
}

func (answerUseCase *answerUseCase) Store(req *domain.AnswerRequest) (domain.Answer, error) {
	answer := domain.Answer{
		QuestionID: req.QuestionID,
		UserID:     req.UserID,
		Answer:     req.Answer,
		File:       req.FileUrl,
		Vote:       0,
		IsCorrect:  false,
	}

	return answerUseCase.answerRepository.Store(answer)
}

func (answerUseCase *answerUseCase) StoreFile(req *domain.AnswerRequestFile) (string, error) {
	return internal.UploadToCloudinary(req.File)
}

func (answerUseCase *answerUseCase) Update(id int, req *domain.AnswerRequest) error {
	getByID, err := answerUseCase.answerRepository.GetByID(id)
	if err != nil {
		return err
	}

	if req.UserID != getByID.UserID {
		return errors.New("You are not authorized to update this answer")
	}

	if req.FileUrl != "" {
		urlResult, _ := internal.GetPublicIDFromURL(getByID.File)
		err := internal.DeleteFromCloudinary(urlResult)
		if err != nil {
			return err
		}
	}

	getByID.Answer = req.Answer
	getByID.File = req.FileUrl

	return answerUseCase.answerRepository.Update(id, &getByID)
}

func (answerUseCase *answerUseCase) Destroy(id int, idLogin uint) error {
	getByID, err := answerUseCase.answerRepository.GetByID(id)
	if err != nil {
		return err
	}

	urlResult, _ := internal.GetPublicIDFromURL(getByID.File)
	err = internal.DeleteFromCloudinary(urlResult)
	if err != nil {
		return err
	}

	if idLogin != getByID.UserID {
		return errors.New("You are not authorized to delete this answer")
	}

	if err := answerUseCase.answerRepository.Destroy(id); err != nil {
		return err
	}

	return nil
}

func (answerUseCase *answerUseCase) MarkAsCorrect(id int, idLogin uint) error {
	getDataAnswer, err := answerUseCase.answerRepository.GetByID(id)
	if err != nil {
		return err
	}

	getDataQuestion, err := answerUseCase.answerRepository.GetQuestionByID(getDataAnswer.QuestionID)
	if err != nil {
		return err
	}

	if idLogin != getDataQuestion.UserID {
		return errors.New("You are not authorized to mark as correct this answer")
	}

	if getDataAnswer.IsCorrect == true {
		return errors.New("Answer already marked as correct")
	}

	getDataAnswer.IsCorrect = true
	return answerUseCase.answerRepository.Update(id, &getDataAnswer)
}

func (answerUseCase *answerUseCase) UpVote(id int, idLogin uint) error {
	getDataAnswer, err := answerUseCase.answerRepository.GetByID(id)
	if err != nil {
		return err
	}

	if idLogin == getDataAnswer.UserID {
		return errors.New("You cannot vote your own answer")
	}

	userVote, _ := answerUseCase.answerRepository.GetUserVote(uint(id), idLogin)

	if userVote.VoteType == 1 {
		return errors.New("You already voted this answer")
	}

	getDataAnswer.Vote++
	if userVote.VoteType == -1 {
		err = answerUseCase.answerRepository.UpdateUserVote(uint(id), idLogin, 1)
	} else {
		err = answerUseCase.answerRepository.AddUserVote(uint(id), idLogin, 1)
	}

	if err != nil {
		return err
	}

	return answerUseCase.answerRepository.Update(id, &getDataAnswer)
}

func (answerUseCase *answerUseCase) DownVote(id int, idLogin uint) error {
	getDataAnswer, err := answerUseCase.answerRepository.GetByID(id)
	if err != nil {
		return err
	}

	if idLogin == getDataAnswer.UserID {
		return errors.New("You cannot vote your own answer")
	}

	userVote, _ := answerUseCase.answerRepository.GetUserVote(uint(id), idLogin)

	if userVote.VoteType == -1 {
		return errors.New("You already voted this answer")
	}

	getDataAnswer.Vote--

	if userVote.VoteType == 1 {
		err = answerUseCase.answerRepository.UpdateUserVote(uint(id), idLogin, -1)
	} else {
		err = answerUseCase.answerRepository.AddUserVote(uint(id), idLogin, -1)
	}

	if err != nil {
		return err
	}

	return answerUseCase.answerRepository.Update(id, &getDataAnswer)
}

func (answerUseCase *answerUseCase) Comment(req *domain.AnswerCommentRequest) error {
	comment := domain.AnswerComment{
		AnswerID: req.AnswerID,
		UserID:   req.UserID,
		Comment:  req.Comment,
	}

	return answerUseCase.answerRepository.Comment(&comment)
}

func (answerUseCase *answerUseCase) DestroyComment(id int, idComment int, idLogin uint) error {
	comment, err := answerUseCase.answerRepository.CommentById(idComment)
	if err != nil {
		return err
	}

	if comment.UserID != idLogin {
		return errors.New("You are not authorized to delete this comment")
	}

	return answerUseCase.answerRepository.DestroyComment(idComment)
}
