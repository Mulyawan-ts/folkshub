package main

import (
	"context"
	"fmt"
	"path/filepath"

	"app/internal/model"
	"app/internal/service"
)

// App struct
type App struct {
	ctx           context.Context
	karyawanSvc   *service.KaryawanService
	absensiSvc    *service.AbsensiService
	cutiSvc       *service.CutiService
	reimburseSvc  *service.ReimburseService
	gajiSvc       *service.GajiService
	jabatanSvc    *service.JabatanService
	departemenSvc *service.DepartemenService
	pelatihanSvc  *service.PelatihanService
}

// NewApp creates a new App application struct storing its SQLite files in the
// current working directory.
func NewApp() *App {
	a, err := newApp(".")
	if err != nil {
		panic(err)
	}
	return a
}

// newApp builds an App with every SQLite file placed inside dbDir.
func newApp(dbDir string) (*App, error) {
	p := func(name string) string { return filepath.Join(dbDir, name+".db") }

	karyawanSvc, err := service.NewKaryawanService(p("karyawan"))
	if err != nil {
		return nil, err
	}
	absensiSvc, err := service.NewAbsensiService(p("absensi"))
	if err != nil {
		return nil, err
	}
	cutiSvc, err := service.NewCutiService(p("cuti"))
	if err != nil {
		return nil, err
	}
	reimburseSvc, err := service.NewReimburseService(p("reimburse"))
	if err != nil {
		return nil, err
	}
	gajiSvc, err := service.NewGajiService(p("gaji"))
	if err != nil {
		return nil, err
	}
	jabatanSvc, err := service.NewJabatanService(p("jabatan"))
	if err != nil {
		return nil, err
	}
	departemenSvc, err := service.NewDepartemenService(p("departemen"))
	if err != nil {
		return nil, err
	}
	pelatihanSvc, err := service.NewPelatihanService(p("pelatihan"))
	if err != nil {
		return nil, err
	}

	return &App{
		karyawanSvc:   karyawanSvc,
		absensiSvc:    absensiSvc,
		cutiSvc:       cutiSvc,
		reimburseSvc:  reimburseSvc,
		gajiSvc:       gajiSvc,
		jabatanSvc:    jabatanSvc,
		departemenSvc: departemenSvc,
		pelatihanSvc:  pelatihanSvc,
	}, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// shutdown menutup semua koneksi SQLite agar file WAL ter-checkpoint sebelum
// proses keluar.
func (a *App) shutdown(_ context.Context) {
	for _, penutup := range []interface{ Close() error }{
		a.karyawanSvc, a.absensiSvc, a.cutiSvc, a.reimburseSvc,
		a.gajiSvc, a.jabatanSvc, a.departemenSvc, a.pelatihanSvc,
	} {
		_ = penutup.Close()
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Karyawan
func (a *App) SimpanKaryawan(nama, email, status string) (*model.Karyawan, error) {
	k := &model.Karyawan{
		Nama:   nama,
		Email:  email,
		Status: status,
	}

	if err := a.karyawanSvc.Simpan(k); err != nil {
		return nil, err
	}

	return k, nil
}

func (a *App) AmbilKaryawan(id int) (*model.Karyawan, error) {
	return a.karyawanSvc.Ambil(id)
}

func (a *App) SemuaKaryawan() ([]*model.Karyawan, error) {
	return a.karyawanSvc.Semua()
}

func (a *App) HapusKaryawan(id int) error {
	return a.karyawanSvc.Hapus(id)
}

func (a *App) UpdateKaryawan(id int, nama, email, status string) (*model.Karyawan, error) {
	return a.karyawanSvc.Update(id, nama, email, status)
}

// Absensi
func (a *App) SimpanAbsensi(karyawanID int, tanggal, checkIn, checkOut, status, keterangan string) (*model.Absensi, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	a2 := &model.Absensi{
		KaryawanId: karyawanID,
		Tanggal:    tanggal,
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		Status:     status,
		Keterangan: keterangan,
	}

	if err := a.absensiSvc.Simpan(a2); err != nil {
		return nil, err
	}

	return a2, nil
}

func (a *App) AmbilAbsensi(id int) (*model.Absensi, error) {
	return a.absensiSvc.Ambil(id)
}

func (a *App) SemuaAbsensi() ([]*model.Absensi, error) {
	return a.absensiSvc.Semua()
}

func (a *App) HapusAbsensi(id int) error {
	return a.absensiSvc.Hapus(id)
}

func (a *App) UpdateAbsensi(id int, karyawanID int, tanggal, checkIn, checkOut, status, keterangan string) (*model.Absensi, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	return a.absensiSvc.Update(id, karyawanID, tanggal, checkIn, checkOut, status, keterangan)
}

// Cuti
func (a *App) SimpanCuti(karyawanID int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}

	c := &model.Cuti{
		KaryawanId:     karyawanID,
		Jenis:          jenis,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		JumlahHari:     jumlahHari,
		Keterangan:     keterangan,
		Status:         status,
	}

	if err := a.cutiSvc.Simpan(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (a *App) AmbilCuti(id int) (*model.Cuti, error) {
	return a.cutiSvc.Ambil(id)
}

func (a *App) SemuaCuti() ([]*model.Cuti, error) {
	return a.cutiSvc.Semua()
}

func (a *App) HapusCuti(id int) error {
	return a.cutiSvc.Hapus(id)
}

func (a *App) UpdateCuti(id int, karyawanID int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	return a.cutiSvc.Update(id, karyawanID, jenis, tanggalMulai, tanggalSelesai, keterangan, status, jumlahHari)
}

// Reimburse
func (a *App) SimpanReimburse(karyawanID int, tanggal, jenis, keterangan, bukti, status string, jumlah int) (*model.Reimburse, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	r := &model.Reimburse{
		KaryawanId: karyawanID,
		Tanggal:    tanggal,
		Jenis:      jenis,
		Jumlah:     jumlah,
		Keterangan: keterangan,
		Bukti:      bukti,
		Status:     status,
	}

	if err := a.reimburseSvc.Simpan(r); err != nil {
		return nil, err
	}

	return r, nil
}

func (a *App) AmbilReimburse(id int) (*model.Reimburse, error) {
	return a.reimburseSvc.Ambil(id)
}

func (a *App) SemuaReimburse() ([]*model.Reimburse, error) {
	return a.reimburseSvc.Semua()
}

func (a *App) HapusReimburse(id int) error {
	return a.reimburseSvc.Hapus(id)
}

func (a *App) UpdateReimburse(id int, karyawanID int, tanggal, jenis, keterangan, bukti, status string, jumlah int) (*model.Reimburse, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	return a.reimburseSvc.Update(id, karyawanID, tanggal, jenis, keterangan, bukti, status, jumlah)
}

// Gaji
func (a *App) SimpanGaji(karyawanID int, bulan string, gajiPokok, tunjangan, potongan int, status string) (*model.Gaji, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	g := &model.Gaji{
		KaryawanId: karyawanID,
		Bulan:      bulan,
		GajiPokok:  gajiPokok,
		Tunjangan:  tunjangan,
		Potongan:   potongan,
		Status:     status,
	}

	if err := a.gajiSvc.Simpan(g); err != nil {
		return nil, err
	}

	return g, nil
}

func (a *App) AmbilGaji(id int) (*model.Gaji, error) {
	return a.gajiSvc.Ambil(id)
}

func (a *App) SemuaGaji() ([]*model.Gaji, error) {
	return a.gajiSvc.Semua()
}

func (a *App) HapusGaji(id int) error {
	return a.gajiSvc.Hapus(id)
}

func (a *App) UpdateGaji(id int, karyawanID int, bulan string, gajiPokok, tunjangan, potongan int, status string) (*model.Gaji, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	return a.gajiSvc.Update(id, karyawanID, bulan, gajiPokok, tunjangan, potongan, status)
}

// Jabatan
func (a *App) SimpanJabatan(kode, nama, departemen, deskripsi string, gajiPokok, tunjangan int) (*model.Jabatan, error) {
	j := &model.Jabatan{
		Kode:       kode,
		Nama:       nama,
		Departemen: departemen,
		GajiPokok:  gajiPokok,
		Tunjangan:  tunjangan,
		Deskripsi:  deskripsi,
	}

	if err := a.jabatanSvc.Simpan(j); err != nil {
		return nil, err
	}

	return j, nil
}

func (a *App) AmbilJabatan(id int) (*model.Jabatan, error) {
	return a.jabatanSvc.Ambil(id)
}

func (a *App) SemuaJabatan() ([]*model.Jabatan, error) {
	return a.jabatanSvc.Semua()
}

func (a *App) HapusJabatan(id int) error {
	return a.jabatanSvc.Hapus(id)
}

func (a *App) UpdateJabatan(id int, kode, nama, departemen, deskripsi string, gajiPokok, tunjangan int) (*model.Jabatan, error) {
	return a.jabatanSvc.Update(id, kode, nama, departemen, deskripsi, gajiPokok, tunjangan)
}

// Departemen
func (a *App) SimpanDepartemen(kode, nama, kepala, deskripsi string) (*model.Departemen, error) {
	d := &model.Departemen{
		Kode:      kode,
		Nama:      nama,
		Kepala:    kepala,
		Deskripsi: deskripsi,
	}

	if err := a.departemenSvc.Simpan(d); err != nil {
		return nil, err
	}

	return d, nil
}

func (a *App) AmbilDepartemen(id int) (*model.Departemen, error) {
	return a.departemenSvc.Ambil(id)
}

func (a *App) SemuaDepartemen() ([]*model.Departemen, error) {
	return a.departemenSvc.Semua()
}

func (a *App) HapusDepartemen(id int) error {
	return a.departemenSvc.Hapus(id)
}

func (a *App) UpdateDepartemen(id int, kode, nama, kepala, deskripsi string) (*model.Departemen, error) {
	return a.departemenSvc.Update(id, kode, nama, kepala, deskripsi)
}

// Pelatihan
func (a *App) SimpanPelatihan(karyawanID int, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status string, biaya int) (*model.Pelatihan, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	p := &model.Pelatihan{
		KaryawanId:     karyawanID,
		Nama:           nama,
		Jenis:          jenis,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		Lokasi:         lokasi,
		Biaya:          biaya,
		Status:         status,
		Sertifikat:     sertifikat,
		Catatan:        catatan,
	}

	if err := a.pelatihanSvc.Simpan(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (a *App) AmbilPelatihan(id int) (*model.Pelatihan, error) {
	return a.pelatihanSvc.Ambil(id)
}

func (a *App) SemuaPelatihan() ([]*model.Pelatihan, error) {
	return a.pelatihanSvc.Semua()
}

func (a *App) HapusPelatihan(id int) error {
	return a.pelatihanSvc.Hapus(id)
}

func (a *App) UpdatePelatihan(id int, karyawanID int, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status string, biaya int) (*model.Pelatihan, error) {
	if _, err := a.karyawanSvc.Ambil(karyawanID); err != nil {
		return nil, err
	}
	return a.pelatihanSvc.Update(id, karyawanID, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status, biaya)
}
