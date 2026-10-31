package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrPembelianNotFound = errors.New("pembelian tidak ditemukan")
	ErrPembelianExists   = errors.New("pembelian sudah ada")
)

var validPembelianJenis = map[string]bool{
	model.JenisATK:       true,
	model.JenisPeralatan: true,
	model.JenisLangganan: true,
	model.JenisJasa:      true,
	model.JenisLainnya:   true,
}

var validPembelianStatus = map[string]bool{
	"diajukan":  true,
	"disetujui": true,
	"dibeli":    true,
	"ditolak":   true,
}

type PembelianService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewPembelianService(dbPath string) (*PembelianService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Pembelian{}); err != nil {
		return nil, err
	}
	return &PembelianService{db: db}, nil
}

func (s *PembelianService) Simpan(p *model.Pembelian) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := validasiPembelian(p); err != nil {
		return err
	}

	p.Total = p.Jumlah * p.HargaSatuan

	if p.ID == 0 {
		return s.db.Create(p).Error
	}

	var existing model.Pembelian
	if err := s.db.First(&existing, p.ID).Error; err == nil {
		return ErrPembelianExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(p).Error
}

func (s *PembelianService) Ambil(id int) (*model.Pembelian, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var p model.Pembelian
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPembelianNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (s *PembelianService) Semua() ([]*model.Pembelian, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Pembelian
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *PembelianService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var p model.Pembelian
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPembelianNotFound
		}
		return err
	}
	return s.db.Delete(&p).Error
}

func (s *PembelianService) Update(id int, p *model.Pembelian) (*model.Pembelian, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p == nil {
		return nil, ErrInvalidInput
	}

	var lama model.Pembelian
	if err := s.db.First(&lama, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPembelianNotFound
		}
		return nil, err
	}

	if err := validasiPembelian(p); err != nil {
		return nil, err
	}

	lama.Tanggal = p.Tanggal
	lama.Barang = p.Barang
	lama.Jenis = p.Jenis
	lama.Jumlah = p.Jumlah
	lama.Satuan = p.Satuan
	lama.HargaSatuan = p.HargaSatuan
	lama.Total = p.Jumlah * p.HargaSatuan
	lama.Vendor = p.Vendor
	lama.Departemen = p.Departemen
	lama.Catatan = p.Catatan
	lama.Status = p.Status

	if err := s.db.Save(&lama).Error; err != nil {
		return nil, err
	}
	return &lama, nil
}

func (s *PembelianService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// validasiPembelian memeriksa isi pembelian tanpa menyentuh database.
func validasiPembelian(p *model.Pembelian) error {
	if p == nil || p.Barang == "" || p.Tanggal == "" {
		return ErrInvalidInput
	}
	if !validDate(p.Tanggal) {
		return ErrInvalidInput
	}
	if !validPembelianJenis[p.Jenis] {
		return ErrInvalidInput
	}
	if !validPembelianStatus[p.Status] {
		return ErrInvalidInput
	}
	if p.Jumlah <= 0 || p.HargaSatuan <= 0 {
		return ErrInvalidInput
	}
	return nil
}
