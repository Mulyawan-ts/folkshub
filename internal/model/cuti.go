package model

type Cuti struct {
	ID          int    `json:"id"`
	KaryawanId  int    `json:"karyawan_id"`
	Jenis       string `json:"jenis"` // tahunan, sakit, izin, melahirkan, lainnya
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	JumlahHari  int    `json:"jumlah_hari"`
	Keterangan  string `json:"keterangan"`
	Status      string `json:"status"` // draft, submitted, approved, rejected, cancelled
	CatatanApproval string `json:"catatan_approval"`
	DisetujuiOleh int   `json:"disetujui_oleh"`
}

func (c *Cuti) TableName() string {
	return "cuti"
}
