package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrReimburseNotFound = errors.New("reimburse tidak ditemukan")
	ErrReimburseExists   = errors.New("reimburse sudah ada")
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

	if r == nil || r.KaryawanId <= 0 || r.Jenis == "" || r.Tanggal == "" {
		return ErrInvalidInput
	}
	if !validDate(r.Tanggal) {
		return ErrInvalidInput
	}

	if r.ID == 0 {
		return s.db.Create(r).Error
	}

	var existing model.Reimburse
	if err := s.db.First(&existing, r.ID).Error; err == nil {
		return ErrReimburseExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(r).Error
}

func (s *ReimburseService) Ambil(id int) (*model.Reimburse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReimburseNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (s *ReimburseService) Semua() ([]*model.Reimburse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Reimburse
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ReimburseService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrReimburseNotFound
		}
		return err
	}
	return s.db.Delete(&r).Error
}

func (s *ReimburseService) Update(id int, karyawanId int, tanggal, jenis, keterangan, bukti, status string, jumlah int) (*model.Reimburse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || jenis == "" || tanggal == "" {
		return nil, ErrInvalidInput
	}

	var r model.Reimburse
	if err := s.db.First(&r, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReimburseNotFound
		}
		return nil, err
	}

	if !validDate(tanggal) {
		return nil, ErrInvalidInput
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

func (s *ReimburseService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}