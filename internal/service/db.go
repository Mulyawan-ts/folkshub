package service

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ErrInvalidInput = errors.New("data tidak valid")

func Open(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(1)
	for _, pragma := range []string{
		"PRAGMA foreign_keys = ON",
		"PRAGMA journal_mode = WAL",
		"PRAGMA busy_timeout = 5000",
	} {
		_ = db.Exec(pragma).Error
	}
	return db, nil
}