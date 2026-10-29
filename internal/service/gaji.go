package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrGajiNotFound = errors.New("gaji tidak ditemukan")
	ErrGajiExists   = errors.New("gaji sudah ada")
)

type GajiService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewGajiService(dbPath string) (*GajiService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Gaji{}); err != nil {
		return nil, err
	}
	return &GajiService{db: db}, nil
}

func (s *GajiService) Simpan(g *model.Gaji) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if g == nil || g.KaryawanId <= 0 || g.Bulan == "" {
		return ErrInvalidInput
	}

	g.Total = g.GajiPokok + g.Tunjangan - g.Potongan

	if g.ID == 0 {
		return s.db.Create(g).Error
	}

	var existing model.Gaji
	if err := s.db.First(&existing, g.ID).Error; err == nil {
		return ErrGajiExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(g).Error
}

func (s *GajiService) Ambil(id int) (*model.Gaji, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGajiNotFound
		}
		return nil, err
	}
	return &g, nil
}

func (s *GajiService) Semua() ([]*model.Gaji, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Gaji
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *GajiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGajiNotFound
		}
		return err
	}
	return s.db.Delete(&g).Error
}

func (s *GajiService) Update(id int, karyawanId int, bulan string, gajiPokok, tunjangan, potongan int, status string) (*model.Gaji, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || bulan == "" {
		return nil, ErrInvalidInput
	}

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGajiNotFound
		}
		return nil, err
	}

	g.KaryawanId = karyawanId
	g.Bulan = bulan
	g.GajiPokok = gajiPokok
	g.Tunjangan = tunjangan
	g.Potongan = potongan
	g.Total = gajiPokok + tunjangan - potongan
	g.Status = status

	if err := s.db.Save(&g).Error; err != nil {
		return nil, err
	}
	return &g, nil
}