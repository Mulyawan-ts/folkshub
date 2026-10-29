package model

import "time"

type Pelatihan struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KaryawanId     int       `json:"karyawan_id" gorm:"not null;index"`
	Nama           string    `json:"nama" gorm:"not null"`
	Jenis          string    `json:"jenis"`
	TanggalMulai   string    `json:"tanggal_mulai" gorm:"not null;index"`
	TanggalSelesai string    `json:"tanggal_selesai"`
	Lokasi         string    `json:"lokasi"`
	Biaya          int       `json:"biaya" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null"` // terjadwal, berlangsung, selesai, dibatalkan
	Sertifikat     string    `json:"sertifikat"`
	Catatan        string    `json:"catatan"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *Pelatihan) TableName() string {
	return "pelatihan"
}