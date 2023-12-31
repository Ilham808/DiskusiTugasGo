package domain

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Question struct {
	*gorm.Model
	UserID      uint
	User        User
	SubjectID   uint
	Subject     Subject
	Question    string
	Description string
	File        string
	Answer      []Answer `gorm:"foreignKey:QuestionID"`
}

type QuestionRequest struct {
	UserID      uint   `json:"-" `
	SubjectID   uint   `validate:"required" json:"subject" form:"subject"`
	Question    string `validate:"required" json:"question" form:"question"`
	Description string `validate:"required" json:"description" form:"description"`
	FileUrl     string `json:"file" form:"file"`
}

type QuestionRequestFile struct {
	File multipart.File
}
type QuestionUseCase interface {
	FetchWithPagination(page, pageSize int) ([]Question, int, error)
	Store(question *QuestionRequest) error
	StoreFile(req *QuestionRequestFile) (string, error)
	DestroyFile(fileUrl string) error
	GetByID(id int) (Question, error)
	Update(id int, question *QuestionRequest) error
	Destroy(id int, idLogin uint) error
	FetchQuestionBySubject(slug string, page, pageSize int) ([]Question, int, error)
	GetIdSubject(slug string) uint
}

type QuestionRepository interface {
	FetchWithPagination(page, pageSize int) ([]Question, int, error)
	Store(question *Question) error
	GetByID(id int) (Question, error)
	GetByIDQuestion(id int) (Question, error)
	Update(id int, question *Question) error
	Destroy(id int) error
	GetIdSubject(slug string) uint
	FetchQuestionBySubject(idSubject, page, pageSize int) ([]Question, int, error)
}
