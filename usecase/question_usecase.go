package usecase

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
)

type questionUseCase struct {
	questionRepository domain.QuestionRepository
}

func NewQuestionUseCase(questionRepository domain.QuestionRepository) domain.QuestionUseCase {
	return &questionUseCase{
		questionRepository: questionRepository,
	}
}

func (q *questionUseCase) FetchWithPagination(page, pageSize int) ([]domain.Question, int, error) {
	questions, totalItems, err := q.questionRepository.FetchWithPagination(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return questions, totalItems, nil
}

func (q *questionUseCase) GetByID(id int) (domain.Question, error) {
	return q.questionRepository.GetByID(id)
}

func (q *questionUseCase) Store(req *domain.QuestionRequest) error {

	question := &domain.Question{
		UserID:      req.UserID,
		SubjectID:   req.SubjectID,
		Question:    req.Question,
		Description: req.Description,
		File:        req.FileUrl,
	}

	return q.questionRepository.Store(question)
}

func (q *questionUseCase) StoreFile(req *domain.QuestionRequestFile) (string, error) {
	return internal.UploadToCloudinary(req.File)
}

func (q *questionUseCase) DestroyFile(fileUrl string) error {
	return internal.DeleteFromCloudinary(fileUrl)
}

func (q *questionUseCase) Update(id int, req *domain.QuestionRequest) error {
	getData, err := q.GetByID(id)
	if err != nil {
		return err
	}

	if req.FileUrl != "" {
		urlResult, _ := internal.GetPublicIDFromURL(getData.File)
		err := internal.DeleteFromCloudinary(urlResult)
		if err != nil {
			return err
		}
	}

	getData.UserID = req.UserID
	getData.SubjectID = req.SubjectID
	getData.Question = req.Question
	getData.Description = req.Description
	getData.File = req.FileUrl

	return q.questionRepository.Update(id, &getData)
}

func (q *questionUseCase) Destroy(id int) error {
	getData, err := q.GetByID(id)
	if err != nil {
		return err
	}
	urlResult, _ := internal.GetPublicIDFromURL(getData.File)
	err = internal.DeleteFromCloudinary(urlResult)
	if err != nil {
		return err
	}
	return q.questionRepository.Destroy(id)
}
