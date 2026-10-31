package service

import (
	"testing"

	"app/internal/model"
)

func TestPelatihanServiceCRUD(t *testing.T) {
	svc, err := NewPelatihanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	p := &model.Pelatihan{KaryawanId: 1, Nama: "Kepemimpinan", Jenis: "internal", TanggalMulai: "2026-11-01", TanggalSelesai: "2026-11-03", Lokasi: "Jakarta", Biaya: 1500000, Status: "terjadwal"}
	if err := svc.Simpan(p); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if p.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	got, err := svc.Ambil(p.ID)
	if err != nil || got.Nama != "Kepemimpinan" || got.Biaya != 1500000 {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(p.ID, 2, "K3", "eksternal", "2026-12-01", "2026-12-02", "Bandung", "CERT-001", "lulus", "selesai", 2000000)
	if err != nil || upd.KaryawanId != 2 || upd.Status != "selesai" || upd.Biaya != 2000000 || upd.Sertifikat != "CERT-001" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := svc.Hapus(p.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(p.ID); err != ErrPelatihanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestPelatihanServiceValidationAndDuplicate(t *testing.T) {
	svc, err := NewPelatihanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	if err := svc.Simpan(&model.Pelatihan{KaryawanId: 0, Nama: "X"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (karyawan 0), got %v", err)
	}
	if err := svc.Simpan(&model.Pelatihan{KaryawanId: 1, Nama: ""}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nama kosong), got %v", err)
	}
	if err := svc.Simpan(nil); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nil), got %v", err)
	}

	p := &model.Pelatihan{KaryawanId: 1, Nama: "Dasar", TanggalMulai: "2026-11-01", Status: "terjadwal"}
	if err := svc.Simpan(p); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if err := svc.Simpan(&model.Pelatihan{ID: p.ID, KaryawanId: 1, Nama: "Dasar 2"}); err != ErrPelatihanExists {
		t.Fatalf("expected exists, got %v", err)
	}

	if _, err := svc.Update(p.ID, 0, "Nama", "", "", "", "", "", "", "", 0); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "Nama", "", "", "", "", "", "", "", 0); err != ErrPelatihanNotFound {
		t.Fatalf("expected not found on update, got %v", err)
	}
	if err := svc.Hapus(999); err != ErrPelatihanNotFound {
		t.Fatalf("expected not found on hapus, got %v", err)
	}
}
