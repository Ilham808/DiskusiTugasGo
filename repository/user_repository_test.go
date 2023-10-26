package repository_test

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	mockDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return mockDB, mock
}

func TestGetUserByEmail(t *testing.T) {
	db, mock := setupTestDB()

	userRepo := repository.NewUserRepository(db)

	tests := []struct {
		name           string
		email          string
		expectedUser   *domain.User
		expectedErr    error
		mockExpectFunc func()
	}{
		{
			name:  "success",
			email: "Ilham Budiawan Sitorus",
			expectedUser: &domain.User{
				Name:  "Ilham Budiawan Sitorus",
				Email: "budiawanilham04@gmail.com",
			},
			expectedErr: nil,
			mockExpectFunc: func() {
				mock.ExpectQuery("SELECT (.+) FROM `users` WHERE `email` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
					WithArgs("budiawanilham04@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"email"}).
						AddRow("budiawanilham04@gmail.com"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpectFunc()

			user, err := userRepo.GetByEmail(tt.email)

			assert.NotNil(t, user)

			assert.Equal(t, tt.expectedUser.Name, user.Name)
			assert.Equal(t, tt.expectedUser.Email, user.Email)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}
