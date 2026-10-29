package model

import "time"

type Cuti struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KaryawanId     int       `json:"karyawan_id" gorm:"not null;index"`
	Jenis          string    `json:"jenis" gorm:"not null"` // tahunan, sakit, izin, melahirkan, lainnya
	TanggalMulai   string    `json:"tanggal_mulai" gorm:"not null;index"`
	TanggalSelesai string    `json:"tanggal_selesai" gorm:"not null"`
	JumlahHari     int       `json:"jumlah_hari" gorm:"not null"`
	Keterangan     string    `json:"keterangan"`
	Status         string    `json:"status" gorm:"not null"` // draft, submitted, approved, rejected, cancelled
	CatatanApproval string   `json:"catatan_approval"`
	DisetujuiOleh  int       `json:"disetujui_oleh" gorm:"index"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (c *Cuti) TableName() string {
	return "cuti"
}