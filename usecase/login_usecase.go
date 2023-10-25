package usecase

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/internal"
	"errors"
)

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (lu *loginUsecase) GetUserByEmail(email string) (domain.User, error) {
	user, err := lu.userRepository.GetByEmail(email)
	if err == nil {
		return *user, nil
	}

	if user == nil {
		return domain.User{}, errors.New("User not found")
	}

	return *user, err
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return internal.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return internal.CreateRefreshToken(user, secret, expiry)
}
