package model

type Gaji struct {
	ID         int    `json:"id"`
	KaryawanId int `json:"karyawan_id"`
	Bulan      string `json:"bulan"`
	GajiPokok  int    `json:"gaji_pokok"`
	Tunjangan  int    `json:"tunjangan"`
	Potongan   int    `json:"potongan"`
	Total      int    `json:"total"`
	Status     string `json:"status"` // draft, diproses, dibayar, ditolak
	TanggalBayar string `json:"tanggal_bayar"`
	Keterangan  string `json:"keterangan"`
}

func (g *Gaji) TableName() string {
	return "gaji"
}
