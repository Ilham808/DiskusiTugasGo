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

func TestSubjectRespository_NewSubjectRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()
	mockDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	subjectRepo := repository.NewSubjectRepository(mockDB)
	assert.NotNil(t, subjectRepo)
}

func TestSubjectRepository_Fetch(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"name", "slug"}).
		AddRow("Testing Subject", "testing-subject")

	expectedSQL := "^SELECT \\* FROM `subjects` WHERE `subjects`.`deleted_at` IS NULL\\s*$"
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	subject, err := repository.NewSubjectRepository(mockDB).Fetch()

	if err != nil {
		t.Fatalf("Error getting subject: %s", err)
	}
	if len(subject) != 1 {
		t.Fatalf("Expected 1 subject, but got %d", len(subject))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %s", err)
	}

	assert.NotNil(t, subject)
}

func TestSubjectRepository_GetByID(t *testing.T) {
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

	subjectRepository := repository.NewSubjectRepository(mockDB)
	subjectID := 1

	expectedSubject := domain.Subject{
		Model: &gorm.Model{
			ID: uint(subjectID),
		},
		Name: "Testing Subject",
		Slug: "testing-subject",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `subjects` WHERE id = ? AND `subjects`.`deleted_at` IS NULL ORDER BY `subjects`.`id` LIMIT 1")).
		WithArgs(subjectID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug"}).
			AddRow(subjectID, "Testing Subject", "testing-subject"))

	subject, err := subjectRepository.GetByID(subjectID)
	assert.NoError(t, err)
	assert.Equal(t, expectedSubject, subject)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSubjectRepository_Store(t *testing.T) {
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

	subjectRepository := repository.NewSubjectRepository(mockDB)

	subject := domain.Subject{
		Name: "Testing Subject",
		Slug: "testing-subject",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `subjects`")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = subjectRepository.Store(&subject)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSubjectRepository_Update(t *testing.T) {
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

	subjectRepository := repository.NewSubjectRepository(mockDB)
	subjectId := 1

	subject := &domain.Subject{
		Model: &gorm.Model{
			ID: uint(subjectId),
		},
		Name: "Updated Subject",
		Slug: "updated-subject",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `subjects`")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = subjectRepository.Update(subject)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestSubjectRepository_Destroy(t *testing.T) {
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

	subjectRepository := repository.NewSubjectRepository(mockDB)
	subjectID := 1

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `subjects` SET `deleted_at`=? WHERE id = ? AND `subjects`.`deleted_at` IS NULL")).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = subjectRepository.Destroy(subjectID)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
