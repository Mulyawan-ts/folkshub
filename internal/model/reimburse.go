package model

type Reimburse struct {
	ID             int    `json:"id"`
	KaryawanID     int    `json:"karyawan_id"`
	Tanggal        string `json:"tanggal"`
	Jenis          string `json:"jenis"`
	Jumlah         int    `json:"jumlah"`
	Keterangan     string `json:"keterangan"`
	Bukti          string `json:"bukti"`
	Status         string `json:"status"` // draft, submitted, approved, rejected, paid
	CatatanApproval string `json:"catatan_approval"`
}

func (r *Reimburse) TableName() string {
	return "reimburse"
}
