package service

import (
	"errors"

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

	if k.ID == 0 {
		return s.db.Create(k).Error
	}

	var existing model.Karyawan
	if err := s.db.First(&existing, k.ID).Error; err == nil {
		return ErrKaryawanExists
	}
	return s.db.Create(k).Error
}

func (s *KaryawanService) Ambil(id int) (*model.Karyawan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		return nil, ErrKaryawanNotFound
	}
	return &k, nil
}

func (s *KaryawanService) Semua() []*model.Karyawan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Karyawan
	s.db.Find(&list)
	return list
}

func (s *KaryawanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		return ErrKaryawanNotFound
	}
	return s.db.Delete(&k).Error
}

func (s *KaryawanService) Update(id int, nama, email, status string) (*model.Karyawan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var k model.Karyawan
	if err := s.db.First(&k, id).Error; err != nil {
		return nil, ErrKaryawanNotFound
	}

	k.Nama = nama
	k.Email = email
	k.Status = status

	if err := s.db.Save(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}
