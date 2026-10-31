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

	if err := svc.Simpan(&model.Gaji{KaryawanId: 0, Bulan: "x"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}