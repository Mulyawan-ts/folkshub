package service

import (
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ErrInvalidInput = errors.New("data tidak valid")

// isoDate is the storage format for all date fields.
const isoDate = "2006-01-02"

// validDate reports whether s is a real calendar date in YYYY-MM-DD form.
func validDate(s string) bool {
	_, err := time.Parse(isoDate, s)
	return err == nil
}

// validOptionalDate allows an empty string, otherwise requires validDate.
func validOptionalDate(s string) bool {
	return s == "" || validDate(s)
}

// validMonth reports whether s is YYYY-MM.
func validMonth(s string) bool {
	_, err := time.Parse("2006-01", s)
	return err == nil
}

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