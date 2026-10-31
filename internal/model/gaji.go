package model

import "time"

type Gaji struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KaryawanId   int       `json:"karyawan_id" gorm:"not null;index"`
	Bulan        string    `json:"bulan" gorm:"not null;index"`
	GajiPokok    int       `json:"gaji_pokok" gorm:"not null"`
	Tunjangan    int       `json:"tunjangan" gorm:"not null"`
	Potongan     int       `json:"potongan" gorm:"not null"`
	Total        int       `json:"total" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null"` // draft, diproses, dibayar, ditolak
	TanggalBayar string    `json:"tanggal_bayar"`
	Keterangan   string    `json:"keterangan"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (g *Gaji) TableName() string {
	return "gaji"
}
