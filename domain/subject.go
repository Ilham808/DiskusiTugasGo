package domain

import "gorm.io/gorm"

type Subject struct {
	*gorm.Model
	Name      string
	Slug      string
	Questions []Question `json:"-" gorm:"foreignkey:SubjectID"`
}

type SubjectRequest struct {
	Name string `validate:"required" json:"name" form:"name"`
}

type SubjectUseCase interface {
	Fetch() ([]Subject, error)
	GetByID(id int) (Subject, error)
	Store(subject *Subject) error
	Update(subject *Subject) error
	Destroy(id int) error
}

type SubjectRepository interface {
	Fetch() ([]Subject, error)
	GetByID(id int) (Subject, error)
	Store(subject *Subject) error
	Update(subject *Subject) error
	Destroy(id int) error
}
