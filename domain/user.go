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
	Store(user *User) error
	GetByID(id int) (User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
}
