package model

import "time"

type Reimburse struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KaryawanId      int       `json:"karyawan_id" gorm:"not null;index"`
	Tanggal         string    `json:"tanggal" gorm:"not null;index"`
	Jenis           string    `json:"jenis" gorm:"not null"`
	Jumlah          int       `json:"jumlah" gorm:"not null"`
	Keterangan      string    `json:"keterangan"`
	Bukti           string    `json:"bukti"`
	Status          string    `json:"status" gorm:"not null"` // draft, submitted, approved, rejected, paid
	CatatanApproval string    `json:"catatan_approval"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (r *Reimburse) TableName() string {
	return "reimburse"
}
