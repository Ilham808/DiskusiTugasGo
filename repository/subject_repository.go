package repository

import (
	"DiskusiTugas/domain"

	"gorm.io/gorm"
)

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) domain.SubjectRepository {
	return &subjectRepository{
		db: db,
	}
}

func (s *subjectRepository) Fetch() ([]domain.Subject, error) {
	var subject []domain.Subject
	if err := s.db.Find(&subject).Error; err != nil {
		return nil, err
	}
	return subject, nil
}

func (s *subjectRepository) GetByID(id int) (domain.Subject, error) {
	var subject domain.Subject
	if err := s.db.Where("id = ?", id).First(&subject).Error; err != nil {
		return domain.Subject{}, err
	}
	return subject, nil
}

func (s *subjectRepository) Store(subject *domain.Subject) error {
	if err := s.db.Create(subject).Error; err != nil {
		return err
	}
	return nil
}

func (s *subjectRepository) Update(subject *domain.Subject) error {
	if err := s.db.Save(subject).Error; err != nil {
		return err
	}
	return nil
}

func (s *subjectRepository) Destroy(id int) error {
	if err := s.db.Where("id = ?", id).Delete(&domain.Subject{}).Error; err != nil {
		return err
	}
	return nil
}
