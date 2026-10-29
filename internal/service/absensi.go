package service

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrAbsensiNotFound = errors.New("absensi tidak ditemukan")
)

type AbsensiService struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewAbsensiService(dbPath string) (*AbsensiService, error) {
	db, err := Open(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Absensi{}); err != nil {
		return nil, err
	}
	return &AbsensiService{db: db}, nil
}

func (s *AbsensiService) Simpan(a *model.Absensi) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if a.ID == 0 {
		return s.db.Create(a).Error
	}

	var existing model.Absensi
	if err := s.db.First(&existing, a.ID).Error; err == nil {
		return errors.New("absensi sudah ada")
	}
	return s.db.Create(a).Error
}

func (s *AbsensiService) Ambil(id int) (*model.Absensi, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		return nil, ErrAbsensiNotFound
	}
	return &a, nil
}

func (s *AbsensiService) Semua() []*model.Absensi {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Absensi
	s.db.Find(&list)
	return list
}

func (s *AbsensiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		return ErrAbsensiNotFound
	}
	return s.db.Delete(&a).Error
}

func (s *AbsensiService) Update(id int, karyawanId int, tanggal, checkIn, checkOut, status, keterangan string) (*model.Absensi, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		return nil, ErrAbsensiNotFound
	}

	a.KaryawanId = karyawanId
	a.Tanggal = tanggal
	a.CheckIn = checkIn
	a.CheckOut = checkOut
	a.Status = status
	a.Keterangan = keterangan

	if err := s.db.Save(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
