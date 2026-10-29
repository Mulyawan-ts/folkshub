package service

import (
	"testing"

	"app/internal/model"
)

func TestKaryawanServiceCRUD(t *testing.T) {
	svc, err := NewKaryawanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	k := &model.Karyawan{Nama: "Budi", Email: "budi@example.com", Status: "tetap"}
	if err := svc.Simpan(k); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if k.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	got, err := svc.Ambil(k.ID)
	if err != nil || got.Nama != "Budi" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(k.ID, "Budi Santoso", "budi2@example.com", "tetap")
	if err != nil || upd.Nama != "Budi Santoso" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := svc.Hapus(k.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(k.ID); err != ErrKaryawanNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestKaryawanServiceValidationAndDuplicate(t *testing.T) {
	svc, err := NewKaryawanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	if err := svc.Simpan(&model.Karyawan{Nama: "", Email: "x@y.com"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}

	k := &model.Karyawan{Nama: "A", Email: "a@b.com"}
	if err := svc.Simpan(k); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	dup := &model.Karyawan{ID: k.ID, Nama: "A2", Email: "a2@b.com"}
	if err := svc.Simpan(dup); err != ErrKaryawanExists {
		t.Fatalf("expected exists, got %v", err)
	}
}