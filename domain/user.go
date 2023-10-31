package domain

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name          string
	Email         string
	Password      string
	Gender        string
	University    string
	Avatar        string
	IsStudent     bool
	Status        string
	Questions     []Question      `json:"-" gorm:"foreignkey:UserID"`
	Answers       []Answer        `json:"-" gorm:"foreignkey:UserID"`
	AnswerComment []AnswerComment `json:"-" gorm:"foreignkey:UserID"`
}

type UserRequest struct {
	Name     string `validate:"required" json:"name" form:"name"`
	Email    string `validate:"required,email" json:"email" form:"email"`
	Password string `validate:"omitempty" json:"password" form:"password"`
	Gender   string `validate:"required" json:"gender" form:"gender"`
	Status   string `validate:"required" json:"status" form:"status"`
}

type UserUseCase interface {
	Fetch() ([]User, error)
	FetchWithPagination(page, pageSize int) ([]User, int, error)
	FecthStudent(page, pageSize int) ([]User, int, error)
	BlockStudent(id int) error
	// CountData() (int64, error)
	GetByEmail(email string) error
	Store(user *User) error
	GetByID(id int) (User, error)
	Update(user *User) error
	Destroy(id int) error
}
type UserRepository interface {
	Fetch() ([]User, error)
	FetchWithPagination(page, pageSize int) ([]User, int, error)
	FecthStudent(page, pageSize int) ([]User, int, error)
	BlockStudent(id int) error
	Store(user *User) error
	GetByID(id int) (User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Destroy(id int) error
}
