package usecase

import "DiskusiTugas/domain"

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) Create(user *domain.User) error {
	return nil
}

func (su *signupUsecase) GetUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}
