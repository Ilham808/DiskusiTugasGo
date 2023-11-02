package repository

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
)

type openAIRepository struct{}

func NewOpenAIRepository() domain.OpenAIRepository {
	return &openAIRepository{}
}

func (o *openAIRepository) GenerateAnswer(req *domain.OpenAIRequest) (*domain.OpenAIRespon, error) {

	ress, err := internal.ResponOpenAI(req.Prompt)
	if err != nil {
		return nil, err
	}

	return &domain.OpenAIRespon{
		Prompt: req.Prompt,
		Answer: ress,
	}, nil
}
