package service

import (
	"app/internal/model"
	"app/internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type KaryawanService struct {
	repo *repository.KaryawanRepository
}

func NewKaryawanService(dbPath string) (*KaryawanService, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	repo := repository.NewKaryawanRepository(db)
	if err := repo.AutoMigrate(); err != nil {
		return nil, err
	}

	return &KaryawanService{repo: repo}, nil
}

func (s *KaryawanService) Simpan(k *model.Karyawan) error {
	return s.repo.Simpan(k)
}

func (s *KaryawanService) Ambil(id int) (*model.Karyawan, error) {
	return s.repo.Ambil(id)
}

func (s *KaryawanService) Semua() []*model.Karyawan {
	list, _ := s.repo.Semua()
	return list
}

func (s *KaryawanService) Hapus(id int) error {
	return s.repo.Hapus(id)
}
