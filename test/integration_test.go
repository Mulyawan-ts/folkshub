package test

import (
	"path/filepath"
	"testing"

	"app/internal/model"
	"app/internal/service"
)

func TestIntegration_KaryawanCutiFlow(t *testing.T) {
	dir := t.TempDir()
	kSvc, err := service.NewKaryawanService(filepath.Join(dir, "karyawan.db"))
	if err != nil {
		t.Fatalf("karyawan service: %v", err)
	}
	cSvc, err := service.NewCutiService(filepath.Join(dir, "cuti.db"))
	if err != nil {
		t.Fatalf("cuti service: %v", err)
	}

	k := &model.Karyawan{Nama: "Andi", Email: "andi@example.com", Status: "tetap"}
	if err := kSvc.Simpan(k); err != nil {
		t.Fatalf("simpan karyawan: %v", err)
	}

	// Valid cuti
	c := &model.Cuti{KaryawanId: k.ID, Jenis: "tahunan", TanggalMulai: "2026-11-01", TanggalSelesai: "2026-11-03", JumlahHari: 3, Status: "submitted", Keterangan: "Liburan"}
	if err := cSvc.Simpan(c); err != nil {
		t.Fatalf("simpan cuti: %v", err)
	}

	// Overlap
	c2 := &model.Cuti{KaryawanId: k.ID, Jenis: "sakit", TanggalMulai: "2026-11-02", TanggalSelesai: "2026-11-02", JumlahHari: 1, Status: "submitted"}
	if err := cSvc.Simpan(c2); err != service.ErrCutiOverlap {
		t.Fatalf("expected overlap, got %v", err)
	}

	// Different karyawan no overlap
	k2 := &model.Karyawan{Nama: "Budi", Email: "budi@example.com", Status: "tetap"}
	if err := kSvc.Simpan(k2); err != nil {
		t.Fatalf("simpan karyawan2: %v", err)
	}
	c3 := &model.Cuti{KaryawanId: k2.ID, Jenis: "tahunan", TanggalMulai: "2026-11-02", TanggalSelesai: "2026-11-02", JumlahHari: 1, Status: "submitted"}
	if err := cSvc.Simpan(c3); err != nil {
		t.Fatalf("simpan cuti k2: %v", err)
	}
}

func TestIntegration_GajiFlow(t *testing.T) {
	dir := t.TempDir()
	kSvc, err := service.NewKaryawanService(filepath.Join(dir, "karyawan.db"))
	if err != nil {
		t.Fatalf("karyawan service: %v", err)
	}
	gSvc, err := service.NewGajiService(filepath.Join(dir, "gaji.db"))
	if err != nil {
		t.Fatalf("gaji service: %v", err)
	}

	k := &model.Karyawan{Nama: "Citra", Email: "citra@example.com", Status: "tetap"}
	if err := kSvc.Simpan(k); err != nil {
		t.Fatalf("simpan karyawan: %v", err)
	}

	g := &model.Gaji{KaryawanId: k.ID, Bulan: "2026-11", GajiPokok: 5000000, Tunjangan: 1000000, Potongan: 500000, Status: "draft"}
	if err := gSvc.Simpan(g); err != nil {
		t.Fatalf("simpan gaji: %v", err)
	}
	if g.Total != 5500000 {
		t.Fatalf("expected total 5500000, got %d", g.Total)
	}

	list, err := gSvc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua gaji: %v / %d", err, len(list))
	}
}