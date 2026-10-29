package model

type Absensi struct {
	ID        int    `json:"id"`
	KaryawanID int   `json:"karyawan_id"`
	Tanggal   string `json:"tanggal"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
	Status    string `json:"status"` // hadir, izin, sakit, alfa, lembur
	Keterangan string `json:"keterangan"`
}

func (a *Absensi) TableName() string {
	return "absensi"
}
