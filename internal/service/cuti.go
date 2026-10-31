package service

import (
	"errors"
	"sync"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrCutiNotFound = errors.New("cuti tidak ditemukan")
	ErrCutiExists   = errors.New("cuti sudah ada")
	ErrCutiOverlap  = errors.New("cuti bertumpuk dengan pengajuan lain")
	ErrCutiInvalid  = errors.New("data cuti tidak valid")
)
var validCutiStatus = map[string]bool{
	"draft":      true,
	"submitted":  true,
	"approved":   true,
	"rejected":   true,
	"cancelled":  true,
}

var validCutiJenis = map[string]bool{
	"tahunan":    true,
	"sakit":      true,
	"izin":       true,
	"melahirkan": true,
	"lainnya":    true,
}

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

	if c == nil || c.KaryawanId <= 0 || c.Jenis == "" || c.TanggalMulai == "" || c.TanggalSelesai == "" {
		return ErrInvalidInput
	}
	if !validDate(c.TanggalMulai) || !validDate(c.TanggalSelesai) {
		return ErrCutiInvalid
	}

	if !validCutiJenis[c.Jenis] {
		return ErrCutiInvalid
	}
	if c.Status == "" {
		c.Status = "draft"
	}
	if !validCutiStatus[c.Status] {
		return ErrCutiInvalid
	}
	if c.TanggalMulai > c.TanggalSelesai {
		return ErrCutiInvalid
	}
	if c.JumlahHari <= 0 {
		return ErrCutiInvalid
	}

	if s.hasOverlap(c.KaryawanId, c.TanggalMulai, c.TanggalSelesai, 0) {
		return ErrCutiOverlap
	}

	if c.ID == 0 {
		return s.db.Create(c).Error
	}

	var existing model.Cuti
	if err := s.db.First(&existing, c.ID).Error; err == nil {
		return ErrCutiExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(c).Error
}

func (s *CutiService) Ambil(id int) (*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCutiNotFound
		}
		return nil, err
	}
	return &c, nil
}

func (s *CutiService) Semua() ([]*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []*model.Cuti
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *CutiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCutiNotFound
		}
		return err
	}
	return s.db.Delete(&c).Error
}

func (s *CutiService) Update(id int, karyawanId int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if karyawanId <= 0 || jenis == "" || tanggalMulai == "" || tanggalSelesai == "" {
		return nil, ErrInvalidInput
	}

	var c model.Cuti
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCutiNotFound
		}
		return nil, err
	}

	if !validCutiJenis[jenis] {
		return nil, ErrCutiInvalid
	}
	if status == "" {
		status = "draft"
	}
	if !validCutiStatus[status] {
		return nil, ErrCutiInvalid
	}
	if !validDate(tanggalMulai) || !validDate(tanggalSelesai) {
		return nil, ErrCutiInvalid
	}
	if tanggalMulai > tanggalSelesai {
		return nil, ErrCutiInvalid
	}
	if jumlahHari <= 0 {
		return nil, ErrCutiInvalid
	}

	if s.hasOverlap(karyawanId, tanggalMulai, tanggalSelesai, id) {
		return nil, ErrCutiOverlap
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
func (s *CutiService) hasOverlap(karyawanId int, start, end string, excludeID int) bool {
	var count int64
	q := s.db.Model(&model.Cuti{}).
		Where("karyawan_id = ? AND status IN ? AND NOT (tanggal_selesai < ? OR tanggal_mulai > ?)",
			karyawanId, []string{"submitted", "approved"}, start, end)
	if excludeID > 0 {
		q = q.Where("id <> ?", excludeID)
	}
	q.Count(&count)
	return count > 0
}

func (s *CutiService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
