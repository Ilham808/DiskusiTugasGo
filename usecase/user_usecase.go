package usecase

import "DiskusiTugas/domain"

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Fetch() ([]domain.User, error) {
	return u.userRepository.Fetch()
}

func (u *userUseCase) FetchWithPagination(page, pageSize int) ([]domain.User, int, error) {
	users, totalItems, err := u.userRepository.FetchWithPagination(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return users, totalItems, nil
}

// func (u *userUseCase) CountData() (int64, error) {
// 	return u.userRepository.CountData()
// }

func (u *userUseCase) Store(user *domain.User) error {
	return u.userRepository.Store(user)
}

func (u *userUseCase) GetByEmail(email string) error {
	_, err := u.userRepository.GetByEmail(email)
	if err == nil {
		return nil
	}

	return err
}

func (u *userUseCase) GetByID(id int) (domain.User, error) {
	return u.userRepository.GetByID(id)
}

func (u *userUseCase) Update(user *domain.User) error {
	return u.userRepository.Update(user)
}

func (u *userUseCase) Destroy(id int) error {
	return nil
}
