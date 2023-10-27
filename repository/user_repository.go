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
	if err := userRepository.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *userRepository) FetchWithPagination(page, pageSize int) ([]domain.User, int, error) {
	var user []domain.User
	var totalItems int64

	if err := userRepository.db.Model(&domain.User{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := userRepository.db.Offset(offset).Limit(pageSize).Find(&user).Error; err != nil {
		return nil, 0, err
	}

	return user, int(totalItems), nil
}

// func (ur *userRepository) CountData() (int64, error) {
// 	var totalItems int64
// 	if err := ur.db.Model(&domain.User{}).Count(&totalItems).Error; err != nil {
// 		return 0, err
// 	}
// 	return totalItems, nil
// }

func (ur *userRepository) Store(user *domain.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
