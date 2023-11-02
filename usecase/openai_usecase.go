package usecase

import "DiskusiTugas/domain"

type openAIUseCase struct {
	openAIRepository domain.OpenAIRepository
}

func NewOpenAIUseCase(openAIRepository domain.OpenAIRepository) domain.OpenAIUseCase {
	return &openAIUseCase{
		openAIRepository: openAIRepository,
	}
}

func (u *openAIUseCase) GenerateAnswer(req *domain.OpenAIRequest) (*domain.OpenAIRespon, error) {
	return u.openAIRepository.GenerateAnswer(req)
}
