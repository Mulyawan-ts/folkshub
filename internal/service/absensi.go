package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrAbsensiNotFound = errors.New("absensi tidak ditemukan")
	ErrAbsensiExists   = errors.New("absensi sudah ada")
)

var validAbsensiStatus = map[string]bool{
	"hadir":  true,
	"izin":   true,
	"sakit":  true,
	"alfa":   true,
	"lembur": true,
}

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

	if a == nil || a.KaryawanId <= 0 || a.Tanggal == "" {
		return ErrInvalidInput
	}
	if !validDate(a.Tanggal) {
		return ErrInvalidInput
	}
	if a.Status != "" && !validAbsensiStatus[a.Status] {
		return ErrInvalidInput
	}

	if a.ID == 0 {
		return s.db.Create(a).Error
	}

	var existing model.Absensi
	if err := s.db.First(&existing, a.ID).Error; err == nil {
		return ErrAbsensiExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(a).Error
}

func (s *AbsensiService) Ambil(id int) (*model.Absensi, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAbsensiNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (s *AbsensiService) Semua() ([]*model.Absensi, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Absensi
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *AbsensiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAbsensiNotFound
		}
		return err
	}
	return s.db.Delete(&a).Error
}

func (s *AbsensiService) Update(id int, karyawanId int, tanggal, checkIn, checkOut, status, keterangan string) (*model.Absensi, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || tanggal == "" {
		return nil, ErrInvalidInput
	}

	var a model.Absensi
	if err := s.db.First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAbsensiNotFound
		}
		return nil, err
	}

	if !validDate(tanggal) {
		return nil, ErrInvalidInput
	}
	if status != "" && !validAbsensiStatus[status] {
		return nil, ErrInvalidInput
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

func (s *AbsensiService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}