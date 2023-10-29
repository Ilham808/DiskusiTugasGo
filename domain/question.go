package domain

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Question struct {
	*gorm.Model
	UserID      uint   `json:"-" `
	SubjectID   uint   `json:"subject" form:"subject"`
	Question    string `json:"question" form:"question"`
	Description string `json:"description" form:"description"`
	File        string `json:"file" form:"file"`
}

type QuestionRequest struct {
	UserID      uint   `json:"user" form:"user"`
	SubjectID   uint   `validate:"required" json:"subject" form:"subject"`
	Question    string `validate:"required" json:"question" form:"question"`
	Description string `validate:"required" json:"description" form:"description"`
	FileUrl     string
}

type QuestionRequestFile struct {
	File multipart.File
}
type QuestionUseCase interface {
	FetchWithPagination(page, pageSize int) ([]Question, int, error)
	Store(question *QuestionRequest) error
	StoreFile(req *QuestionRequestFile) (string, error)
}

type QuestionRepository interface {
	FetchWithPagination(page, pageSize int) ([]Question, int, error)
	Store(question *Question) error
}
