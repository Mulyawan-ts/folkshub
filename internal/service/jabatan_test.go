package service

import (
	"testing"

	"app/internal/model"
)

func TestJabatanServiceCRUD(t *testing.T) {
	svc, err := NewJabatanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	j := &model.Jabatan{Kode: "MGR", Nama: "Manager", Departemen: "HRD", GajiPokok: 8000000, Tunjangan: 2000000}
	if err := svc.Simpan(j); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if j.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	got, err := svc.Ambil(j.ID)
	if err != nil || got.Nama != "Manager" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(j.ID, "MGR", "Senior Manager", "OPS", "Naik level", 10000000, 2500000)
	if err != nil || upd.Nama != "Senior Manager" || upd.GajiPokok != 10000000 || upd.Tunjangan != 2500000 {
		t.Fatalf("update: %v / %+v", err, upd)
	}
	if upd.Deskripsi != "Naik level" {
		t.Fatalf("expected deskripsi updated, got %q", upd.Deskripsi)
	}

	if err := svc.Hapus(j.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(j.ID); err != ErrJabatanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestJabatanServiceValidationAndDuplicate(t *testing.T) {
	svc, err := NewJabatanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	if err := svc.Simpan(&model.Jabatan{Kode: "", Nama: "X"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (kode kosong), got %v", err)
	}
	if err := svc.Simpan(&model.Jabatan{Kode: "X", Nama: ""}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nama kosong), got %v", err)
	}
	if err := svc.Simpan(nil); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nil), got %v", err)
	}

	j := &model.Jabatan{Kode: "STF", Nama: "Staff"}
	if err := svc.Simpan(j); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if err := svc.Simpan(&model.Jabatan{ID: j.ID, Kode: "STF", Nama: "Staff 2"}); err != ErrJabatanExists {
		t.Fatalf("expected exists, got %v", err)
	}

	if _, err := svc.Update(j.ID, "", "Nama", "", "", 0, 0); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, "K", "N", "", "", 0, 0); err != ErrJabatanNotFound {
		t.Fatalf("expected not found on update, got %v", err)
	}
	if err := svc.Hapus(999); err != ErrJabatanNotFound {
		t.Fatalf("expected not found on hapus, got %v", err)
	}
}
