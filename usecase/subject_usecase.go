package usecase

import "DiskusiTugas/domain"

type subjectUseCase struct {
	subjectRepository domain.SubjectRepository
}

func NewSubjectUseCase(subjectRepository domain.SubjectRepository) domain.SubjectUseCase {
	return &subjectUseCase{
		subjectRepository: subjectRepository,
	}
}

func (s *subjectUseCase) Fetch() ([]domain.Subject, error) {
	return s.subjectRepository.Fetch()
}

func (s *subjectUseCase) GetByID(id int) (domain.Subject, error) {
	return s.subjectRepository.GetByID(id)
}

func (s *subjectUseCase) Store(subject *domain.Subject) error {
	return s.subjectRepository.Store(subject)
}

func (s *subjectUseCase) Update(subject *domain.Subject) error {
	return s.subjectRepository.Update(subject)
}

func (s *subjectUseCase) Destroy(id int) error {
	return s.subjectRepository.Destroy(id)
}
