package config

import (
	"DiskusiTugas/domain"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Subject{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Question{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.Answer{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&domain.AnswerComment{}); err != nil {
		return err
	}
	//Tambahkan struct lainnya jika ada yang ingin dimigrate

	return nil
}
