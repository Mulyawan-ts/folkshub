package service

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrGajiNotFound = errors.New("gaji tidak ditemukan")
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

	g.Total = g.GajiPokok + g.Tunjangan - g.Potongan

	if g.ID == 0 {
		return s.db.Create(g).Error
	}

	var existing model.Gaji
	if err := s.db.First(&existing, g.ID).Error; err == nil {
		return errors.New("gaji sudah ada")
	}
	return s.db.Create(g).Error
}

func (s *GajiService) Ambil(id int) (*model.Gaji, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		return nil, ErrGajiNotFound
	}
	return &g, nil
}

func (s *GajiService) Semua() []*model.Gaji {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Gaji
	s.db.Find(&list)
	return list
}

func (s *GajiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		return ErrGajiNotFound
	}
	return s.db.Delete(&g).Error
}

func (s *GajiService) Update(id int, karyawanId int, bulan string, gajiPokok, tunjangan, potongan int, status string) (*model.Gaji, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var g model.Gaji
	if err := s.db.First(&g, id).Error; err != nil {
		return nil, ErrGajiNotFound
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
