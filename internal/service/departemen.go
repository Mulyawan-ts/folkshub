package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrDepartemenNotFound = errors.New("departemen tidak ditemukan")
	ErrDepartemenExists   = errors.New("departemen sudah ada")
)

type DepartemenService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewDepartemenService(dbPath string) (*DepartemenService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Departemen{}); err != nil {
		return nil, err
	}
	return &DepartemenService{db: db}, nil
}

func (s *DepartemenService) Simpan(d *model.Departemen) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if d == nil || d.Kode == "" || d.Nama == "" {
		return ErrInvalidInput
	}

	if d.ID == 0 {
		return s.db.Create(d).Error
	}

	var existing model.Departemen
	if err := s.db.First(&existing, d.ID).Error; err == nil {
		return ErrDepartemenExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(d).Error
}

func (s *DepartemenService) Ambil(id int) (*model.Departemen, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var d model.Departemen
	if err := s.db.First(&d, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDepartemenNotFound
		}
		return nil, err
	}
	return &d, nil
}

func (s *DepartemenService) Semua() ([]*model.Departemen, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Departemen
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *DepartemenService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var d model.Departemen
	if err := s.db.First(&d, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrDepartemenNotFound
		}
		return err
	}
	return s.db.Delete(&d).Error
}

func (s *DepartemenService) Update(id int, kode, nama, kepala, deskripsi string) (*model.Departemen, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if kode == "" || nama == "" {
		return nil, ErrInvalidInput
	}

	var d model.Departemen
	if err := s.db.First(&d, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDepartemenNotFound
		}
		return nil, err
	}

	d.Kode = kode
	d.Nama = nama
	d.Kepala = kepala
	d.Deskripsi = deskripsi

	if err := s.db.Save(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (s *DepartemenService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}