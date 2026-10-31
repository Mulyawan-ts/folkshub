package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrKaryawanNotFound = errors.New("karyawan tidak ditemukan")
	ErrKaryawanExists   = errors.New("karyawan sudah terdaftar")
)

type KaryawanService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewKaryawanService(dbPath string) (*KaryawanService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Karyawan{}); err != nil {
		return nil, err
	}
	return &KaryawanService{db: db}, nil
}

func (s *KaryawanService) Simpan(k *model.Karyawan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if k == nil || k.Nama == "" || k.Email == "" {
		return ErrInvalidInput
	}

	if k.ID == 0 {
		return s.db.Create(k).Error
	}

	var existing model.Karyawan
	if err := s.db.First(&existing, k.ID).Error; err == nil {
		return ErrKaryawanExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(k).Error
}

func (s *KaryawanService) Ambil(id int) (*model.Karyawan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrKaryawanNotFound
		}
		return nil, err
	}
	return &k, nil
}

func (s *KaryawanService) Semua() ([]*model.Karyawan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Karyawan
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *KaryawanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrKaryawanNotFound
		}
		return err
	}
	return s.db.Delete(&k).Error
}

func (s *KaryawanService) Update(id int, nama, email, status string) (*model.Karyawan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if nama == "" || email == "" {
		return nil, ErrInvalidInput
	}

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrKaryawanNotFound
		}
		return nil, err
	}

	k.Nama = nama
	k.Email = email
	k.Status = status

	if err := s.db.Save(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (s *KaryawanService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}