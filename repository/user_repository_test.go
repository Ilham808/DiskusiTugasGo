package repository_test

import (
	"DiskusiTugas/domain"
	"DiskusiTugas/repository"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserReporsitory_NewUserRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()
	mockDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	userRepo := repository.NewUserRepository(mockDB)

	assert.NotNil(t, userRepo)
}

func TestUserRepository_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	rows := sqlmock.NewRows([]string{"name", "email", "gender", "status", "avatar"}).
		AddRow("Testing User", "testinguser@gmail.com", "L", "active", "avatar.jpg")

	expectedSQL := "^SELECT \\* FROM `users` WHERE `users`.`deleted_at` IS NULL\\s*$"
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	users, err := repository.NewUserRepository(mockDB).Fetch()

	if err != nil {
		t.Fatalf("Error getting users: %s", err)
	}
	if len(users) != 1 {
		t.Fatalf("Expected 1 user, but got %d", len(users))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %s", err)
	}

	assert.Equal(t, users[0].Name, "Testing User")
	assert.Equal(t, users[0].Email, "testinguser@gmail.com")
	assert.Equal(t, users[0].Gender, "L")
	assert.Equal(t, users[0].Status, "active")
	assert.Equal(t, users[0].Avatar, "avatar.jpg")

}

func TestUserRepository_FetchWithPagination(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	countQuery := "^SELECT count\\(\\*\\) FROM `users` WHERE is_student = \\? AND `users`.`deleted_at` IS NULL$"
	mock.ExpectQuery(countQuery).
		WithArgs(false).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	rows := sqlmock.NewRows([]string{"name", "email", "gender", "status", "avatar"}).
		AddRow("Testing User 1", "testuser1@gmail.com", "L", "active", "avatar.jpg").
		AddRow("Testing User 2", "testuser2@gmail.com", "P", "active", "avatar.jpg")
	selectQuery := regexp.QuoteMeta("SELECT id,name, email, gender, status, avatar FROM `users` WHERE is_student = ? AND `users`.`deleted_at` IS NULL LIMIT 10")
	mock.ExpectQuery(selectQuery).
		WillReturnRows(rows)

	users, totalItems, err := repository.NewUserRepository(mockDB).FetchWithPagination(1, 10)
	if err != nil {
		t.Fatalf("Error getting users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("Expected 1 user, but got %d", len(users))
	}

	if totalItems != 2 {
		t.Fatalf("Expected total items to be 2, but got %d", totalItems)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %s", err)
	}
}

func TestUserReporsitory_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)

	user := &domain.User{
		Name:       "Test User",
		Email:      "testuser@gmail.com",
		Password:   "123456",
		Gender:     "P",
		University: "ITB",
		Avatar:     "avatar.jpg",
		IsStudent:  false,
		Status:     "active",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = userRepository.Store(user)

	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUserReporsitory_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)
	userID := 1
	expectedUser := domain.User{
		Model: &gorm.Model{
			ID: uint(userID),
		},
		Name:       "Test User",
		Email:      "testuser@gmail.com",
		Password:   "123456",
		Gender:     "P",
		University: "ITB",
		Avatar:     "avatar.jpg",
		IsStudent:  false,
		Status:     "active",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WithArgs("testuser@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "gender", "university", "avatar", "is_student", "status"}).
			AddRow(userID, expectedUser.Name, expectedUser.Email, expectedUser.Password, expectedUser.Gender, expectedUser.University, expectedUser.Avatar, expectedUser.IsStudent, expectedUser.Status))

	user, err := userRepository.GetByEmail("testuser@gmail.com")
	actualUser := *user
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestUserReporsitory_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)
	userID := 1

	expectedUser := domain.User{
		Model: &gorm.Model{
			ID: uint(userID),
		},
		Name:       "Test User",
		Email:      "testuser@gmail.com",
		Password:   "123456",
		Gender:     "P",
		University: "ITB",
		Avatar:     "avatar.jpg",
		IsStudent:  false,
		Status:     "active",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "gender", "university", "avatar", "is_student", "status"}).
			AddRow(userID, expectedUser.Name, expectedUser.Email, expectedUser.Password, expectedUser.Gender, expectedUser.University, expectedUser.Avatar, expectedUser.IsStudent, expectedUser.Status))

	user, err := userRepository.GetByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUserReporsitory_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)
	userID := 1

	serToUpdate := &domain.User{
		Model: &gorm.Model{
			ID: uint(userID),
		},
		Name:       "Updated Name",
		Email:      "updated@example.com",
		Password:   "updatedpassword",
		Gender:     "P",
		University: "ITB",
		Avatar:     "updatedavatar.jpg",
		IsStudent:  false,
		Status:     "block",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET")).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err = userRepository.Update(serToUpdate)

	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestUserReporsitory_Destroy(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = userRepository.Destroy(1)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestUserRepository_FetchStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	countQuery := "^SELECT count\\(\\*\\) FROM `users` WHERE is_student = \\? AND `users`.`deleted_at` IS NULL$"
	mock.ExpectQuery(countQuery).
		WithArgs(true).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	rows := sqlmock.NewRows([]string{"name", "email", "gender", "status", "avatar"}).
		AddRow("Testing User 1", "testuser1@gmail.com", "L", "active", "avatar.jpg").
		AddRow("Testing User 2", "testuser2@gmail.com", "P", "active", "avatar.jpg")
	selectQuery := regexp.QuoteMeta("SELECT id,name, email, gender, status, avatar FROM `users` WHERE is_student = ? AND `users`.`deleted_at` IS NULL LIMIT 10")
	mock.ExpectQuery(selectQuery).
		WillReturnRows(rows)

	users, totalItems, err := repository.NewUserRepository(mockDB).FecthStudent(1, 10)
	if err != nil {
		t.Fatalf("Error getting users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("Expected 1 user, but got %d", len(users))
	}

	if totalItems != 2 {
		t.Fatalf("Expected total items to be 2, but got %d", totalItems)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %s", err)
	}
}

func TestUserRepository_BlockStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()

	mockDB, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if gormErr != nil {
		t.Fatalf("Error creating GORM instance: %s", gormErr)
	}

	userRepository := repository.NewUserRepository(mockDB)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "status"}).AddRow(1, "Test User", "test@example.com", "active"))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET")).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = userRepository.BlockStudent(1)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}
