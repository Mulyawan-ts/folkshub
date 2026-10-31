package model

import "time"

// Pembelian adalah pembelian operasional kantor (ATK, peralatan, langganan),
// berbeda dari Reimburse yang mengganti uang karyawan.
type Pembelian struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Tanggal     string    `json:"tanggal" gorm:"not null;index"`
	Barang      string    `json:"barang" gorm:"not null"`
	Jenis       string    `json:"jenis" gorm:"not null"`
	Jumlah      int       `json:"jumlah" gorm:"not null"`
	Satuan      string    `json:"satuan"`
	HargaSatuan int       `json:"harga_satuan" gorm:"not null"`
	Total       int       `json:"total" gorm:"not null"`
	Vendor      string    `json:"vendor"`
	Departemen  string    `json:"departemen"`
	Catatan     string    `json:"catatan"`
	Status      string    `json:"status" gorm:"not null"` // diajukan, disetujui, dibeli, ditolak
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

const (
	JenisATK       = "atk"
	JenisPeralatan = "peralatan"
	JenisLangganan = "langganan"
	JenisJasa      = "jasa"
	JenisLainnya   = "lainnya"
)

func (p *Pembelian) TableName() string {
	return "pembelian"
}
