package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrJabatanNotFound = errors.New("jabatan tidak ditemukan")
	ErrJabatanExists   = errors.New("jabatan sudah ada")
)

type JabatanService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewJabatanService(dbPath string) (*JabatanService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Jabatan{}); err != nil {
		return nil, err
	}
	return &JabatanService{db: db}, nil
}

func (s *JabatanService) Simpan(j *model.Jabatan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if j == nil || j.Kode == "" || j.Nama == "" {
		return ErrInvalidInput
	}

	if j.ID == 0 {
		return s.db.Create(j).Error
	}

	var existing model.Jabatan
	if err := s.db.First(&existing, j.ID).Error; err == nil {
		return ErrJabatanExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(j).Error
}

func (s *JabatanService) Ambil(id int) (*model.Jabatan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var j model.Jabatan
	if err := s.db.First(&j, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrJabatanNotFound
		}
		return nil, err
	}
	return &j, nil
}

func (s *JabatanService) Semua() ([]*model.Jabatan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Jabatan
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *JabatanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var j model.Jabatan
	if err := s.db.First(&j, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrJabatanNotFound
		}
		return err
	}
	return s.db.Delete(&j).Error
}

func (s *JabatanService) Update(id int, kode, nama, departemen, deskripsi string, gajiPokok, tunjangan int) (*model.Jabatan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if kode == "" || nama == "" {
		return nil, ErrInvalidInput
	}

	var j model.Jabatan
	if err := s.db.First(&j, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrJabatanNotFound
		}
		return nil, err
	}

	j.Kode = kode
	j.Nama = nama
	j.Departemen = departemen
	j.GajiPokok = gajiPokok
	j.Tunjangan = tunjangan
	j.Deskripsi = deskripsi

	if err := s.db.Save(&j).Error; err != nil {
		return nil, err
	}
	return &j, nil
}