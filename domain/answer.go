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

type UserVote struct {
	*gorm.Model
	UserID   uint
	AnswerID uint
	VoteType int
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

type AnswerUsecase interface {
	Store(answer *AnswerRequest) (Answer, error)
	StoreFile(req *AnswerRequestFile) (string, error)
	Update(id int, answer *AnswerRequest) error
	Destroy(id int, idLogin uint) error
	MarkAsCorrect(id int, idLogin uint) error
	UpVote(id int, idLogin uint) error
	DownVote(id int, idLogin uint) error
	Comment(req *AnswerCommentRequest) error
	DestroyComment(id int, idComment int, idLogin uint) error
}

type AnswerRepository interface {
	Store(answer Answer) (Answer, error)
	Update(id int, answer *Answer) error
	Destroy(id int) error
	GetByID(id int) (Answer, error)
	GetQuestionByID(id uint) (Question, error)
	GetUserVote(id uint, idLogin uint) (UserVote, error)
	AddUserVote(id uint, idLogin uint, voteType int) error
	UpdateUserVote(id uint, idLogin uint, voteType int) error
	Comment(comment *AnswerComment) error
	DestroyComment(idComment int) error
	CommentById(idComment int) (AnswerComment, error)
}
