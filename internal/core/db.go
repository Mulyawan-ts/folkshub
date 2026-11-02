package core

import (
	"app/internal/service"
	"gorm.io/gorm"
)

func Open(path string) (*gorm.DB, error) {
	return service.Open(path)
}
