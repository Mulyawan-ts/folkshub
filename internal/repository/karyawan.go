package repository

import (
	"errors"

	"app/internal/model"
	"gorm.io/gorm"
)

var (
	ErrKaryawanNotFound = errors.New("karyawan tidak ditemukan")
	ErrKaryawanExists   = errors.New("karyawan sudah terdaftar")
)

type KaryawanRepository struct {
	db *gorm.DB
}

func NewKaryawanRepository(db *gorm.DB) *KaryawanRepository {
	return &KaryawanRepository{db: db}
}

func (r *KaryawanRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.Karyawan{})
}

func (r *KaryawanRepository) Simpan(k *model.Karyawan) error {
	if k.ID == 0 {
		var count int64
		r.db.Model(&model.Karyawan{}).Where("email = ?", k.Email).Count(&count)
		if count > 0 {
			return ErrKaryawanExists
		}
		return r.db.Create(k).Error
	}

	var existing model.Karyawan
	if err := r.db.First(&existing, k.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrKaryawanNotFound
		}
		return err
	}

	return r.db.Model(&model.Karyawan{}).Where("id = ?", k.ID).Updates(k).Error
}

func (r *KaryawanRepository) Ambil(id int) (*model.Karyawan, error) {
	var k model.Karyawan
	if err := r.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrKaryawanNotFound
		}
		return nil, err
	}
	return &k, nil
}

func (r *KaryawanRepository) Semua() ([]*model.Karyawan, error) {
	var list []*model.Karyawan
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *KaryawanRepository) Hapus(id int) error {
	var k model.Karyawan
	if err := r.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrKaryawanNotFound
		}
		return err
	}
	return r.db.Delete(&k).Error
}
