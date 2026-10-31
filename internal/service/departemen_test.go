package service

import (
	"testing"

	"app/internal/model"
)

func TestDepartemenServiceCRUD(t *testing.T) {
	svc, err := NewDepartemenService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	d := &model.Departemen{Kode: "HRD", Nama: "Sumber Daya Manusia", Kepala: "Andi"}
	if err := svc.Simpan(d); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if d.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	got, err := svc.Ambil(d.ID)
	if err != nil || got.Nama != "Sumber Daya Manusia" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(d.ID, "HRD", "Human Resources", "Budi", "Urusan SDM")
	if err != nil || upd.Nama != "Human Resources" || upd.Kepala != "Budi" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := svc.Hapus(d.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(d.ID); err != ErrDepartemenNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestDepartemenServiceValidationAndDuplicate(t *testing.T) {
	svc, err := NewDepartemenService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	if err := svc.Simpan(&model.Departemen{Kode: "", Nama: "X"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (kode kosong), got %v", err)
	}
	if err := svc.Simpan(&model.Departemen{Kode: "X", Nama: ""}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nama kosong), got %v", err)
	}
	if err := svc.Simpan(nil); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nil), got %v", err)
	}

	d := &model.Departemen{Kode: "OPS", Nama: "Operasional"}
	if err := svc.Simpan(d); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if err := svc.Simpan(&model.Departemen{ID: d.ID, Kode: "OPS", Nama: "Operasional 2"}); err != ErrDepartemenExists {
		t.Fatalf("expected exists, got %v", err)
	}

	if _, err := svc.Update(d.ID, "", "Nama", "", ""); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, "K", "N", "", ""); err != ErrDepartemenNotFound {
		t.Fatalf("expected not found on update, got %v", err)
	}
	if err := svc.Hapus(999); err != ErrDepartemenNotFound {
		t.Fatalf("expected not found on hapus, got %v", err)
	}
}
