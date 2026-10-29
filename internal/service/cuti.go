package service

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrCutiNotFound = errors.New("cuti tidak ditemukan")
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

	if c.ID == 0 {
		return s.db.Create(c).Error
	}

	var existing model.Cuti
	if err := s.db.First(&existing, c.ID).Error; err == nil {
		return errors.New("cuti sudah ada")
	}
	return s.db.Create(c).Error
}

func (s *CutiService) Ambil(id int) (*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, ErrCutiNotFound
	}
	return &c, nil
}

func (s *CutiService) Semua() []*model.Cuti {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Cuti
	s.db.Find(&list)
	return list
}

func (s *CutiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		return ErrCutiNotFound
	}
	return s.db.Delete(&c).Error
}

func (s *CutiService) Update(id int, karyawanId int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, ErrCutiNotFound
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
