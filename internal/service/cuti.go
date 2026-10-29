package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrCutiNotFound = errors.New("cuti tidak ditemukan")
	ErrCutiExists   = errors.New("cuti sudah ada")
)

type CutiService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewCutiService(dbPath string) (*CutiService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Cuti{}); err != nil {
		return nil, err
	}
	return &CutiService{db: db}, nil
}

func (s *CutiService) Simpan(c *model.Cuti) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if c == nil || c.KaryawanId <= 0 || c.Jenis == "" || c.TanggalMulai == "" || c.TanggalSelesai == "" {
		return ErrInvalidInput
	}

	if c.ID == 0 {
		return s.db.Create(c).Error
	}

	var existing model.Cuti
	if err := s.db.First(&existing, c.ID).Error; err == nil {
		return ErrCutiExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(c).Error
}

func (s *CutiService) Ambil(id int) (*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCutiNotFound
		}
		return nil, err
	}
	return &c, nil
}

func (s *CutiService) Semua() ([]*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Cuti
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *CutiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCutiNotFound
		}
		return err
	}
	return s.db.Delete(&c).Error
}

func (s *CutiService) Update(id int, karyawanId int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || jenis == "" || tanggalMulai == "" || tanggalSelesai == "" {
		return nil, ErrInvalidInput
	}

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCutiNotFound
		}
		return nil, err
	}

	c.KaryawanId = karyawanId
	c.Jenis = jenis
	c.TanggalMulai = tanggalMulai
	c.TanggalSelesai = tanggalSelesai
	c.JumlahHari = jumlahHari
	c.Keterangan = keterangan
	c.Status = status

	if err := s.db.Save(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}