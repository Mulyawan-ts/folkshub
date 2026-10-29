package service

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ErrInvalidInput = errors.New("data tidak valid")

func Open(path string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}