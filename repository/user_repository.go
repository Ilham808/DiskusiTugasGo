package repository

import (
	"DiskusiTugas/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
func (userRepository *userRepository) Fetch() ([]domain.User, error) {
	var user []domain.User
	return user, nil
}

func (ur *userRepository) Store(user *domain.User) error {
	return nil
}

func (ur *userRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	return user, nil
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	return user, nil
}
