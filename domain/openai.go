package domain

type OpenAIRequest struct {
	Prompt string `validate:"required" json:"prompt" form:"prompt"`
}

type OpenAIRespon struct {
	Prompt string `json:"prompt"`
	Answer string `json:"answer"`
}

type OpenAIUseCase interface {
	GenerateAnswer(req *OpenAIRequest) (*OpenAIRespon, error)
}

type OpenAIRepository interface {
	GenerateAnswer(req *OpenAIRequest) (*OpenAIRespon, error)
}
