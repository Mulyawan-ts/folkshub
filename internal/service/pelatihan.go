package service

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrPelatihanNotFound = errors.New("pelatihan tidak ditemukan")
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

	if p.ID == 0 {
		return s.db.Create(p).Error
	}

	var existing model.Pelatihan
	if err := s.db.First(&existing, p.ID).Error; err == nil {
		return errors.New("pelatihan sudah ada")
	}
	return s.db.Create(p).Error
}

func (s *PelatihanService) Ambil(id int) (*model.Pelatihan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		return nil, ErrPelatihanNotFound
	}
	return &p, nil
}

func (s *PelatihanService) Semua() []*model.Pelatihan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Pelatihan
	s.db.Find(&list)
	return list
}

func (s *PelatihanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		return ErrPelatihanNotFound
	}
	return s.db.Delete(&p).Error
}

func (s *PelatihanService) Update(id int, karyawanId int, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status string, biaya int) (*model.Pelatihan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var p model.Pelatihan
	if err := s.db.First(&p, id).Error; err != nil {
		return nil, ErrPelatihanNotFound
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
