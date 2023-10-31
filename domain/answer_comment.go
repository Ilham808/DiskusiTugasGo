package domain

import "gorm.io/gorm"

type AnswerComment struct {
	*gorm.Model
	AnswerID uint
	UserID   uint
	User     User
	Comment  string
}

type AnswerCommentRequest struct {
	UserID   uint
	AnswerID uint
	Comment  string `validate:"required" json:"comment" form:"comment"`
}
