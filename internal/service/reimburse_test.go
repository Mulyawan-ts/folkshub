package service

import (
	"testing"

	"app/internal/model"
)

func TestReimburseServiceCRUD(t *testing.T) {
	svc, err := NewReimburseService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	r := &model.Reimburse{KaryawanId: 1, Tanggal: "2026-10-15", Jenis: "transport", Jumlah: 250000, Keterangan: "Perjalanan dinas", Status: "submitted"}
	if err := svc.Simpan(r); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if r.ID == 0 {
		t.Fatalf("expected id assigned")
	}

	got, err := svc.Ambil(r.ID)
	if err != nil || got.Jumlah != 250000 || got.Jenis != "transport" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 1 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	upd, err := svc.Update(r.ID, 2, "2026-10-20", "kesehatan", "Cek gigi", "bukti.pdf", "approved", 500000)
	if err != nil || upd.KaryawanId != 2 || upd.Jenis != "kesehatan" || upd.Jumlah != 500000 || upd.Status != "approved" || upd.Bukti != "bukti.pdf" {
		t.Fatalf("update: %v / %+v", err, upd)
	}

	if err := svc.Hapus(r.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	if _, err := svc.Ambil(r.ID); err != ErrReimburseNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestReimburseServiceValidationAndDuplicate(t *testing.T) {
	svc, err := NewReimburseService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	if err := svc.Simpan(&model.Reimburse{KaryawanId: 0, Tanggal: "2026-10-01", Jenis: "x"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (karyawan 0), got %v", err)
	}
	if err := svc.Simpan(&model.Reimburse{KaryawanId: 1, Tanggal: "2026-10-01", Jenis: ""}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (jenis kosong), got %v", err)
	}
	if err := svc.Simpan(&model.Reimburse{KaryawanId: 1, Tanggal: "", Jenis: "x"}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (tanggal kosong), got %v", err)
	}
	if err := svc.Simpan(nil); err != ErrInvalidInput {
		t.Fatalf("expected invalid input (nil), got %v", err)
	}

	r := &model.Reimburse{KaryawanId: 1, Tanggal: "2026-10-01", Jenis: "transport", Status: "draft"}
	if err := svc.Simpan(r); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if err := svc.Simpan(&model.Reimburse{ID: r.ID, KaryawanId: 1, Tanggal: "2026-10-02", Jenis: "transport"}); err != ErrReimburseExists {
		t.Fatalf("expected exists, got %v", err)
	}

	if _, err := svc.Update(r.ID, 0, "", "", "", "", "", 0); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "2026-10-01", "x", "", "", "", 0); err != ErrReimburseNotFound {
		t.Fatalf("expected not found on update, got %v", err)
	}
	if err := svc.Hapus(999); err != ErrReimburseNotFound {
		t.Fatalf("expected not found on hapus, got %v", err)
	}
}

func TestReimburseServiceDateValidation(t *testing.T) {
	svc, err := NewReimburseService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, tanggal := range []string{"15-10-2026", "2026-10-32", "2026-02-30", "kemarin"} {
		if err := svc.Simpan(&model.Reimburse{KaryawanId: 1, Tanggal: tanggal, Jenis: "transport", Status: "submitted"}); err != ErrInvalidInput {
			t.Fatalf("tanggal %q: expected invalid input, got %v", tanggal, err)
		}
	}

	r := &model.Reimburse{KaryawanId: 1, Tanggal: "2026-10-15", Jenis: "transport", Status: "draft"}
	if err := svc.Simpan(r); err != nil {
		t.Fatalf("tanggal valid ditolak: %v", err)
	}
	if _, err := svc.Update(r.ID, 1, "2026-10-32", "transport", "", "", "", 0); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, 1, "2026-10-32", "transport", "", "", "", 0); err != ErrReimburseNotFound {
		t.Fatalf("expected not found before date check on update, got %v", err)
	}
}
