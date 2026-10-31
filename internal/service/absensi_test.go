package service

import (
	"testing"

	"app/internal/model"
)

func TestAbsensiServiceCRUD(t *testing.T) {
	svc, err := NewAbsensiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	a := &model.Absensi{KaryawanId: 1, Tanggal: "2026-10-01", Status: "hadir"}
	if err := svc.Simpan(a); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if a.ID == 0 {
		t.Fatalf("expected id")
	}

	got, err := svc.Ambil(a.ID)
	if err != nil || got.KaryawanId != 1 {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(a.ID, 2, "2026-10-02", "08:00", "17:00", "hadir", "ok")
	if err != nil || upd.KaryawanId != 2 {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := svc.Hapus(a.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(a.ID); err != ErrAbsensiNotFound {
		t.Fatalf("expected not found, got %v", err)
	}

	if err := svc.Simpan(&model.Absensi{KaryawanId: 0, Tanggal: "x"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestAbsensiServiceDateValidation(t *testing.T) {
	svc, err := NewAbsensiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, tanggal := range []string{"01-10-2026", "2026-13-01", "2026-02-30", "kemarin"} {
		if err := svc.Simpan(&model.Absensi{KaryawanId: 1, Tanggal: tanggal, Status: "hadir"}); err != ErrInvalidInput {
			t.Fatalf("tanggal %q: expected invalid input, got %v", tanggal, err)
		}
	}

	a := &model.Absensi{KaryawanId: 1, Tanggal: "2026-10-01", Status: "hadir"}
	if err := svc.Simpan(a); err != nil {
		t.Fatalf("tanggal valid ditolak: %v", err)
	}
	if _, err := svc.Update(a.ID, 1, "2026-10-32", "", "", "", ""); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "2026-10-32", "", "", "", ""); err != ErrAbsensiNotFound {
		t.Fatalf("expected not found before date check on update, got %v", err)
	}
}

func TestAbsensiServiceStatusValidation(t *testing.T) {
	svc, err := NewAbsensiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, status := range []string{"hadir2", "HADIR", "cuti", "lembur "} {
		err := svc.Simpan(&model.Absensi{KaryawanId: 1, Tanggal: "2026-10-01", Status: status})
		if err != ErrInvalidInput {
			t.Fatalf("status %q: expected invalid input, got %v", status, err)
		}
	}

	for _, status := range []string{"hadir", "izin", "sakit", "alfa", "lembur"} {
		a := &model.Absensi{KaryawanId: 1, Tanggal: "2026-10-01", Status: status}
		if err := svc.Simpan(a); err != nil {
			t.Fatalf("status %q ditolak: %v", status, err)
		}
	}

	if err := svc.Simpan(&model.Absensi{KaryawanId: 1, Tanggal: "2026-10-02"}); err != nil {
		t.Fatalf("status kosong ditolak: %v", err)
	}

	a := &model.Absensi{KaryawanId: 1, Tanggal: "2026-10-03", Status: "hadir"}
	if err := svc.Simpan(a); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if _, err := svc.Update(a.ID, 1, "2026-10-03", "", "", "cuti", ""); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "2026-10-03", "", "", "cuti", ""); err != ErrAbsensiNotFound {
		t.Fatalf("expected not found before status check on update, got %v", err)
	}
	upd, err := svc.Update(a.ID, 1, "2026-10-03", "08:00", "17:00", "lembur", "ok")
	if err != nil || upd.Status != "lembur" {
		t.Fatalf("update status valid: %v / %+v", err, upd)
	}
}

func TestGajiServiceTotal(t *testing.T) {
	svc, err := NewGajiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	g := &model.Gaji{KaryawanId: 1, Bulan: "2026-10", GajiPokok: 1000, Tunjangan: 200, Potongan: 50, Status: "draft"}
	if err := svc.Simpan(g); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if g.Total != 1150 {
		t.Fatalf("expected total 1150, got %d", g.Total)
	}

	upd, err := svc.Update(g.ID, 1, "2026-10", 2000, 300, 100, "diproses")
	if err != nil || upd.Total != 2200 {
		t.Fatalf("update total: %v / %d", err, upd.Total)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 || list[0].ID != g.ID {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	for _, bulan := range []string{"2026-13", "11-2026", "2026-11-01", "november"} {
		if err := svc.Simpan(&model.Gaji{KaryawanId: 1, Bulan: bulan}); err != ErrInvalidInput {
			t.Fatalf("bulan %q: expected invalid input, got %v", bulan, err)
		}
	}

	if err := svc.Simpan(&model.Gaji{KaryawanId: 0, Bulan: "x"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestGajiServiceStatusValidation(t *testing.T) {
	svc, err := NewGajiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, status := range []string{"draft2", "DRAFT", "bayar", "diproses "} {
		err := svc.Simpan(&model.Gaji{KaryawanId: 1, Bulan: "2026-11", Status: status})
		if err != ErrInvalidInput {
			t.Fatalf("status %q: expected invalid input, got %v", status, err)
		}
	}

	for _, status := range []string{"draft", "diproses", "dibayar", "ditolak"} {
		g := &model.Gaji{KaryawanId: 1, Bulan: "2026-11", Status: status}
		if err := svc.Simpan(g); err != nil {
			t.Fatalf("status %q ditolak: %v", status, err)
		}
	}

	if err := svc.Simpan(&model.Gaji{KaryawanId: 1, Bulan: "2026-12"}); err != nil {
		t.Fatalf("status kosong ditolak: %v", err)
	}

	g := &model.Gaji{KaryawanId: 1, Bulan: "2027-01", Status: "draft"}
	if err := svc.Simpan(g); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if _, err := svc.Update(g.ID, 1, "2027-01", 0, 0, 0, "bayar"); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "2027-01", 0, 0, 0, "bayar"); err != ErrGajiNotFound {
		t.Fatalf("expected not found before status check on update, got %v", err)
	}
	upd, err := svc.Update(g.ID, 1, "2027-01", 1000, 0, 0, "dibayar")
	if err != nil || upd.Status != "dibayar" || upd.Total != 1000 {
		t.Fatalf("update status valid: %v / %+v", err, upd)
	}
}