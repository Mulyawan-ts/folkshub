package main

import (
	"path/filepath"
	"testing"

	"app/internal/model"
	"app/internal/service"
)

// newTestApp builds an App backed by per-test SQLite files in t.TempDir,
// mirroring NewApp without the fixed working-directory paths.
func newTestApp(t *testing.T) *App {
	t.Helper()
	dir := t.TempDir()
	p := func(name string) string { return filepath.Join(dir, name+".db") }

	ks, err := service.NewKaryawanService(p("karyawan"))
	if err != nil {
		t.Fatalf("karyawan svc: %v", err)
	}
	t.Cleanup(func() { _ = ks.Close() })
	as, err := service.NewAbsensiService(p("absensi"))
	if err != nil {
		t.Fatalf("absensi svc: %v", err)
	}
	t.Cleanup(func() { _ = as.Close() })
	cs, err := service.NewCutiService(p("cuti"))
	if err != nil {
		t.Fatalf("cuti svc: %v", err)
	}
	t.Cleanup(func() { _ = cs.Close() })
	rs, err := service.NewReimburseService(p("reimburse"))
	if err != nil {
		t.Fatalf("reimburse svc: %v", err)
	}
	t.Cleanup(func() { _ = rs.Close() })
	gs, err := service.NewGajiService(p("gaji"))
	if err != nil {
		t.Fatalf("gaji svc: %v", err)
	}
	t.Cleanup(func() { _ = gs.Close() })
	js, err := service.NewJabatanService(p("jabatan"))
	if err != nil {
		t.Fatalf("jabatan svc: %v", err)
	}
	t.Cleanup(func() { _ = js.Close() })
	ds, err := service.NewDepartemenService(p("departemen"))
	if err != nil {
		t.Fatalf("departemen svc: %v", err)
	}
	t.Cleanup(func() { _ = ds.Close() })
	ps, err := service.NewPelatihanService(p("pelatihan"))
	if err != nil {
		t.Fatalf("pelatihan svc: %v", err)
	}
	t.Cleanup(func() { _ = ps.Close() })

	return &App{
		karyawanSvc:   ks,
		absensiSvc:    as,
		cutiSvc:       cs,
		reimburseSvc:  rs,
		gajiSvc:       gs,
		jabatanSvc:    js,
		departemenSvc: ds,
		pelatihanSvc:  ps,
	}
}

func simpanKaryawan(t *testing.T, a *App, nama, email string) *model.Karyawan {
	t.Helper()
	k, err := a.SimpanKaryawan(nama, email, "tetap")
	if err != nil {
		t.Fatalf("simpan karyawan %q: %v", nama, err)
	}
	return k
}

func TestAppKaryawan(t *testing.T) {
	a := newTestApp(t)

	k, err := a.SimpanKaryawan("Andi", "andi@example.com", "tetap")
	if err != nil || k.ID == 0 {
		t.Fatalf("simpan: %v / %+v", err, k)
	}

	got, err := a.AmbilKaryawan(k.ID)
	if err != nil || got.Nama != "Andi" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaKaryawan()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateKaryawan(k.ID, "Andi S", "andi2@example.com", "kontrak")
	if err != nil || upd.Nama != "Andi S" || upd.Status != "kontrak" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := a.HapusKaryawan(k.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilKaryawan(k.ID); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}

	if _, err := a.SimpanKaryawan("", "x@y.com", ""); err != service.ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestAppAbsensi(t *testing.T) {
	a := newTestApp(t)

	if _, err := a.SimpanAbsensi(99, "2026-10-01", "", "", "", ""); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found, got %v", err)
	}
	k1 := simpanKaryawan(t, a, "Andi", "andi@example.com")
	k2 := simpanKaryawan(t, a, "Budi", "budi@example.com")

	ab, err := a.SimpanAbsensi(k1.ID, "2026-10-01", "08:00", "17:00", "hadir", "ok")
	if err != nil || ab.ID == 0 || ab.CheckIn != "08:00" || ab.KaryawanId != k1.ID {
		t.Fatalf("simpan: %v / %+v", err, ab)
	}

	got, err := a.AmbilAbsensi(ab.ID)
	if err != nil || got.KaryawanId != k1.ID {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaAbsensi()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateAbsensi(ab.ID, k2.ID, "2026-10-02", "09:00", "18:00", "hadir", "telat")
	if err != nil || upd.KaryawanId != k2.ID || upd.CheckIn != "09:00" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if _, err := a.UpdateAbsensi(ab.ID, 99, "2026-10-02", "", "", "", ""); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found on update, got %v", err)
	}

	if err := a.HapusAbsensi(ab.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilAbsensi(ab.ID); err != service.ErrAbsensiNotFound {
		t.Fatalf("expected not found, got %v", err)
	}

	if _, err := a.SimpanAbsensi(k1.ID, "", "", "", "", ""); err != service.ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestAppCutiValidatesKaryawan(t *testing.T) {
	a := newTestApp(t)

	// karyawan tidak ada -> app layer menolak sebelum service cuti
	if _, err := a.SimpanCuti(99, "tahunan", "2026-11-01", "2026-11-02", "", "submitted", 2); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found, got %v", err)
	}

	k, err := a.SimpanKaryawan("Budi", "budi@example.com", "tetap")
	if err != nil {
		t.Fatalf("simpan karyawan: %v", err)
	}

	c, err := a.SimpanCuti(k.ID, "tahunan", "2026-11-01", "2026-11-03", "Liburan", "submitted", 3)
	if err != nil || c.ID == 0 || c.KaryawanId != k.ID {
		t.Fatalf("simpan cuti: %v / %+v", err, c)
	}

	got, err := a.AmbilCuti(c.ID)
	if err != nil || got.Jenis != "tahunan" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaCuti()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	if _, err := a.UpdateCuti(c.ID, 99, "tahunan", "2026-11-05", "2026-11-06", "", "submitted", 2); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found on update, got %v", err)
	}

	upd, err := a.UpdateCuti(c.ID, k.ID, "sakit", "2026-11-05", "2026-11-06", "Demam", "approved", 2)
	if err != nil || upd.Jenis != "sakit" || upd.Status != "approved" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	// overlap setelah update
	if _, err := a.SimpanCuti(k.ID, "izin", "2026-11-06", "2026-11-07", "", "submitted", 2); err != service.ErrCutiOverlap {
		t.Fatalf("expected overlap, got %v", err)
	}

	if err := a.HapusCuti(c.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilCuti(c.ID); err != service.ErrCutiNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppReimburse(t *testing.T) {
	a := newTestApp(t)

	if _, err := a.SimpanReimburse(99, "2026-10-15", "transport", "", "", "", 1); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found, got %v", err)
	}
	k1 := simpanKaryawan(t, a, "Andi", "andi@example.com")
	k2 := simpanKaryawan(t, a, "Budi", "budi@example.com")

	r, err := a.SimpanReimburse(k1.ID, "2026-10-15", "transport", "Dinas", "bukti.pdf", "submitted", 250000)
	if err != nil || r.ID == 0 || r.Jumlah != 250000 || r.KaryawanId != k1.ID {
		t.Fatalf("simpan: %v / %+v", err, r)
	}

	got, err := a.AmbilReimburse(r.ID)
	if err != nil || got.Jenis != "transport" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaReimburse()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateReimburse(r.ID, k2.ID, "2026-10-20", "kesehatan", "Gigi", "b2.pdf", "approved", 500000)
	if err != nil || upd.Jenis != "kesehatan" || upd.Jumlah != 500000 || upd.KaryawanId != k2.ID {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if _, err := a.UpdateReimburse(r.ID, 99, "2026-10-20", "kesehatan", "", "", "", 1); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found on update, got %v", err)
	}

	if err := a.HapusReimburse(r.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilReimburse(r.ID); err != service.ErrReimburseNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppGaji(t *testing.T) {
	a := newTestApp(t)

	if _, err := a.SimpanGaji(99, "2026-11", 1, 0, 0, "draft"); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found, got %v", err)
	}
	k1 := simpanKaryawan(t, a, "Andi", "andi@example.com")

	g, err := a.SimpanGaji(k1.ID, "2026-11", 5000000, 1000000, 500000, "draft")
	if err != nil || g.ID == 0 || g.Total != 5500000 || g.KaryawanId != k1.ID {
		t.Fatalf("simpan: %v / %+v", err, g)
	}

	got, err := a.AmbilGaji(g.ID)
	if err != nil || got.Bulan != "2026-11" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaGaji()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateGaji(g.ID, k1.ID, "2026-11", 2000000, 300, 100, "diproses")
	if err != nil || upd.Total != 2000200 {
		t.Fatalf("update: %v / %d", err, upd.Total)
	}

	if _, err := a.UpdateGaji(g.ID, 99, "2026-11", 1, 0, 0, ""); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found on update, got %v", err)
	}

	if err := a.HapusGaji(g.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilGaji(g.ID); err != service.ErrGajiNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppJabatan(t *testing.T) {
	a := newTestApp(t)

	j, err := a.SimpanJabatan("MGR", "Manager", "HRD", "Level manajerial", 8000000, 2000000)
	if err != nil || j.ID == 0 || j.GajiPokok != 8000000 {
		t.Fatalf("simpan: %v / %+v", err, j)
	}

	got, err := a.AmbilJabatan(j.ID)
	if err != nil || got.Kode != "MGR" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaJabatan()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateJabatan(j.ID, "MGR", "Senior Manager", "OPS", "Naik", 10000000, 2500000)
	if err != nil || upd.Nama != "Senior Manager" || upd.Tunjangan != 2500000 {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := a.HapusJabatan(j.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilJabatan(j.ID); err != service.ErrJabatanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppDepartemen(t *testing.T) {
	a := newTestApp(t)

	d, err := a.SimpanDepartemen("HRD", "Sumber Daya Manusia", "Andi", "Urusan SDM")
	if err != nil || d.ID == 0 || d.Kode != "HRD" {
		t.Fatalf("simpan: %v / %+v", err, d)
	}

	got, err := a.AmbilDepartemen(d.ID)
	if err != nil || got.Nama != "Sumber Daya Manusia" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaDepartemen()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdateDepartemen(d.ID, "HRD", "Human Resources", "Budi", "SDM")
	if err != nil || upd.Nama != "Human Resources" || upd.Kepala != "Budi" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := a.HapusDepartemen(d.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilDepartemen(d.ID); err != service.ErrDepartemenNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppPelatihan(t *testing.T) {
	a := newTestApp(t)

	if _, err := a.SimpanPelatihan(99, "X", "", "2026-11-01", "", "", "", "", "", 0); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found, got %v", err)
	}
	k1 := simpanKaryawan(t, a, "Andi", "andi@example.com")
	k2 := simpanKaryawan(t, a, "Budi", "budi@example.com")

	p, err := a.SimpanPelatihan(k1.ID, "Kepemimpinan", "internal", "2026-11-01", "2026-11-03", "Jakarta", "CERT-1", "lulus", "terjadwal", 1500000)
	if err != nil || p.ID == 0 || p.Biaya != 1500000 || p.KaryawanId != k1.ID {
		t.Fatalf("simpan: %v / %+v", err, p)
	}

	got, err := a.AmbilPelatihan(p.ID)
	if err != nil || got.Nama != "Kepemimpinan" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := a.SemuaPelatihan()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := a.UpdatePelatihan(p.ID, k2.ID, "K3", "eksternal", "2026-12-01", "2026-12-02", "Bandung", "CERT-2", "selesai", "selesai", 2000000)
	if err != nil || upd.KaryawanId != k2.ID || upd.Biaya != 2000000 || upd.Sertifikat != "CERT-2" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if _, err := a.UpdatePelatihan(p.ID, 99, "K3", "", "2026-12-01", "", "", "", "", "", 0); err != service.ErrKaryawanNotFound {
		t.Fatalf("expected karyawan not found on update, got %v", err)
	}

	if err := a.HapusPelatihan(p.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := a.AmbilPelatihan(p.ID); err != service.ErrPelatihanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestAppGreet(t *testing.T) {
	a := newTestApp(t)
	if got := a.Greet("Andi"); got != "Hello Andi, It's show time!" {
		t.Fatalf("greet: %q", got)
	}
}
