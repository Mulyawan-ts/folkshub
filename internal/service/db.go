package service

import (
	"errors"
	"time"
)

var ErrInvalidInput = errors.New("data tidak valid")

const isoDate = "2006-01-02"

func validDate(s string) bool {
	_, err := time.Parse(isoDate, s)
	return err == nil
}

func validOptionalDate(s string) bool {
	return s == "" || validDate(s)
}

func validMonth(s string) bool {
	_, err := time.Parse("2006-01", s)
	return err == nil
}
