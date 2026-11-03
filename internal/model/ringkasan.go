package model

// RingkasanDashboard adalah DTO agregat untuk halaman dashboard HRIS desktop.
type RingkasanDashboard struct {
	TotalKaryawan       int `json:"total_karyawan"`
	KaryawanTetap       int `json:"karyawan_tetap"`
	KaryawanOutsourcing  int `json:"karyawan_outsourcing"`
	KaryawanMagang      int `json:"karyawan_magang"`

	TotalAbsensi        int `json:"total_absensi"`
	TotalCuti           int `json:"total_cuti"`
	CutiPending         int `json:"cuti_pending"`

	TotalReimburse      int `json:"total_reimburse"`
	ReimbursePending    int `json:"reimburse_pending"`

	TotalGaji           int `json:"total_gaji"`
	TotalJabatan        int `json:"total_jabatan"`
	TotalDepartemen     int `json:"total_departemen"`
	TotalPelatihan      int `json:"total_pelatihan"`
}
