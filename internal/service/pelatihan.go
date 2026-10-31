package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrPelatihanNotFound = errors.New("pelatihan tidak ditemukan")
	ErrPelatihanExists   = errors.New("pelatihan sudah ada")
)

type PelatihanService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewPelatihanService(dbPath string) (*PelatihanService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Pelatihan{}); err != nil {
		return nil, err
	}
	return &PelatihanService{db: db}, nil
}

func (s *PelatihanService) Simpan(p *model.Pelatihan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p == nil || p.KaryawanId <= 0 || p.Nama == "" {
		return ErrInvalidInput
	}
	if !validDate(p.TanggalMulai) || !validOptionalDate(p.TanggalSelesai) {
		return ErrInvalidInput
	}

	if p.ID == 0 {
		return s.db.Create(p).Error
	}

	var existing model.Pelatihan
	if err := s.db.First(&existing, p.ID).Error; err == nil {
		return ErrPelatihanExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(p).Error
}

func (s *PelatihanService) Ambil(id int) (*model.Pelatihan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPelatihanNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (s *PelatihanService) Semua() ([]*model.Pelatihan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Pelatihan
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *PelatihanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPelatihanNotFound
		}
		return err
	}
	return s.db.Delete(&p).Error
}

func (s *PelatihanService) Update(id int, karyawanId int, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status string, biaya int) (*model.Pelatihan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || nama == "" {
		return nil, ErrInvalidInput
	}

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPelatihanNotFound
		}
		return nil, err
	}

	if !validDate(tanggalMulai) || !validOptionalDate(tanggalSelesai) {
		return nil, ErrInvalidInput
	}

	p.KaryawanId = karyawanId
	p.Nama = nama
	p.Jenis = jenis
	p.TanggalMulai = tanggalMulai
	p.TanggalSelesai = tanggalSelesai
	p.Lokasi = lokasi
	p.Biaya = biaya
	p.Sertifikat = sertifikat
	p.Catatan = catatan
	p.Status = status

	if err := s.db.Save(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *PelatihanService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}