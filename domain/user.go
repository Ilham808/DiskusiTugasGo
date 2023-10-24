package domain

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Gender     string `json:"gender" form:"gender"`
	University string `json:"university" form:"university"`
	Avatar     string `json:"avatar" form:"avatar"`
	IsStudent  bool   `json:"is_student" form:"is_student"`
	Status     string `json:"status" form:"status"`
}

type UserRepository interface {
	Fetch() ([]User, error)
	Store(user *User) error
	GetByID(id int) (User, error)
	GetByEmail(email string) (*User, error)
}
