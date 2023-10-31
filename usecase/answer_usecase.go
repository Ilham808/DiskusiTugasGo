package usecase

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
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

func (answerUseCase *answerUseCase) Destroy(id int) error {
	return answerUseCase.answerRepository.Destroy(id)
}
