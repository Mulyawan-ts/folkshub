package service

import (
	"fmt"
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

	if err := svc.Simpan(&model.Karyawan{Nama: "", Email: "x@y.com", Status: model.StatusTetap}); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}

	k := &model.Karyawan{Nama: "A", Email: "a@b.com", Status: model.StatusTetap}
	if err := svc.Simpan(k); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	dup := &model.Karyawan{ID: k.ID, Nama: "A2", Email: "a2@b.com", Status: model.StatusTetap}
	if err := svc.Simpan(dup); err != ErrKaryawanExists {
		t.Fatalf("expected exists, got %v", err)
	}
}

func TestKaryawanServiceStatusValidation(t *testing.T) {
	svc, err := NewKaryawanService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, status := range []string{"", "honorer", "TETAP", "kontrak"} {
		err := svc.Simpan(&model.Karyawan{Nama: "A", Email: "a@b.com", Status: status})
		if err != ErrInvalidInput {
			t.Fatalf("status %q: expected invalid input, got %v", status, err)
		}
	}

	for i, status := range []string{model.StatusTetap, model.StatusOutsourcing, model.StatusMagang} {
		k := &model.Karyawan{Nama: "K", Email: fmt.Sprintf("k%d@b.com", i), Status: status}
		if err := svc.Simpan(k); err != nil {
			t.Fatalf("status %q ditolak: %v", status, err)
		}
	}

	k := &model.Karyawan{Nama: "B", Email: "b@b.com", Status: model.StatusTetap}
	if err := svc.Simpan(k); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if _, err := svc.Update(k.ID, "B", "b@b.com", "honorer"); err != ErrInvalidInput {
		t.Fatalf("expected invalid input on update, got %v", err)
	}
	if _, err := svc.Update(999, "B", "b@b.com", "honorer"); err != ErrKaryawanNotFound {
		t.Fatalf("expected not found before status check on update, got %v", err)
	}
	upd, err := svc.Update(k.ID, "B", "b@b.com", model.StatusMagang)
	if err != nil || upd.Status != model.StatusMagang {
		t.Fatalf("update status valid: %v / %+v", err, upd)
	}
}
