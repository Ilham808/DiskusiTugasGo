package repository_test

import (
	"DiskusiTugas/repository"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestQuestionRepository_NewQuestionRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating mock database: %s", err)
	}
	defer db.Close()
	mockDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	questionRepo := repository.NewQuestionRepository(mockDB)
	assert.NotNil(t, questionRepo)
}

func TestQuestionRepository_FetchWithPagination(t *testing.T) {
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

	countQuery := "^SELECT count\\(\\*\\) FROM `questions` WHERE `questions`.`deleted_at` IS NULL$"
	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, user_id, subject_id, question, description, file FROM `questions` WHERE `questions`.`deleted_at` IS NULL LIMIT 10").WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "subject_id", "question", "description", "file"}))
	questionRepo := repository.NewQuestionRepository(mockDB)
	questions, _, err := questionRepo.FetchWithPagination(1, 10)

	if err != nil {
		t.Errorf("Error getting questions: %v", err)
	}

	for _, question := range questions {
		assert.Nil(t, question.DeletedAt)
	}
	for _, question := range questions {
		assert.NotNil(t, question.UserID)
		assert.NotNil(t, question.SubjectID)
	}

	for _, question := range questions {
		assert.NotNil(t, question.Question)
		assert.NotNil(t, question.Description)
	}

	for _, question := range questions {
		if question.File != "nil" {
			assert.NotNil(t, question.File)
		}
	}

}
