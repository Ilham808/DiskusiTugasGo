package usecase

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
)

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) Store(user *domain.User) error {
	return su.userRepository.Store(user)
}

func (su *signupUsecase) GetUserByEmail(email string) error {
	_, err := su.userRepository.GetByEmail(email)
	if err == nil {
		return nil
	}

	return err
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return internal.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return internal.CreateRefreshToken(user, secret, expiry)
}
