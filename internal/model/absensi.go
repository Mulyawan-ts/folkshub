package model

import "time"

type Absensi struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KaryawanId int       `json:"karyawan_id" gorm:"not null;index"`
	Tanggal    string    `json:"tanggal" gorm:"not null;index"`
	CheckIn    string    `json:"check_in"`
	CheckOut   string    `json:"check_out"`
	Status     string    `json:"status" gorm:"not null"` // hadir, izin, sakit, alfa, lembur
	Keterangan string    `json:"keterangan"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *Absensi) TableName() string {
	return "absensi"
}