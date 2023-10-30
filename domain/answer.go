package domain

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Answer struct {
	*gorm.Model
	UserID     uint
	User       User
	QuestionID uint
	Answer     string
	File       string
	Vote       int
	IsCorrect  bool
	Comment    []AnswerComment `gorm:"foreignKey:AnswerID"`
}

type AnswerRequest struct {
	UserID     uint
	QuestionID uint   `json:"question_id" form:"question_id"`
	Answer     string `validate:"required" json:"answer" form:"answer"`
	FileUrl    string `json:"file" form:"file"`
	Vote       int    `json:"vote" form:"vote"`
	IsCorrect  bool   `json:"is_correct" form:"is_correct"`
}

type AnswerRequestFile struct {
	File multipart.File
}
