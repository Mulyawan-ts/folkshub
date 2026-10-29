package model

type Pelatihan struct {
	ID           int    `json:"id"`
	KaryawanID   int    `json:"karyawan_id"`
	Nama         string `json:"nama"`
	Jenis        string `json:"jenis"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	Lokasi       string `json:"lokasi"`
	Biaya        int    `json:"biaya"`
	Status       string `json:"status"` // terjadwal, berlangsung, selesai, dibatalkan
	Sertifikat   string `json:"sertifikat"`
	Catatan      string `json:"catatan"`
}

func (p *Pelatihan) TableName() string {
	return "pelatihan"
}
