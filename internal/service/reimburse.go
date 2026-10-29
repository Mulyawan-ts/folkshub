package service

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrReimburseNotFound = errors.New("reimburse tidak ditemukan")
)

type ReimburseService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewReimburseService(dbPath string) (*ReimburseService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Reimburse{}); err != nil {
		return nil, err
	}
	return &ReimburseService{db: db}, nil
}

func (s *ReimburseService) Simpan(r *model.Reimburse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r.ID == 0 {
		return s.db.Create(r).Error
	}

	var existing model.Reimburse
	if err := s.db.First(&existing, r.ID).Error; err == nil {
		return errors.New("reimburse sudah ada")
	}
	return s.db.Create(r).Error
}

func (s *ReimburseService) Ambil(id int) (*model.Reimburse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		return nil, ErrReimburseNotFound
	}
	return &r, nil
}

func (s *ReimburseService) Semua() []*model.Reimburse {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Reimburse
	s.db.Find(&list)
	return list
}

func (s *ReimburseService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		return ErrReimburseNotFound
	}
	return s.db.Delete(&r).Error
}

func (s *ReimburseService) Update(id int, karyawanId int, tanggal, jenis, keterangan, bukti, status string, jumlah int) (*model.Reimburse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		return nil, ErrReimburseNotFound
	}

	r.KaryawanId = karyawanId
	r.Tanggal = tanggal
	r.Jenis = jenis
	r.Jumlah = jumlah
	r.Keterangan = keterangan
	r.Bukti = bukti
	r.Status = status

	if err := s.db.Save(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}
