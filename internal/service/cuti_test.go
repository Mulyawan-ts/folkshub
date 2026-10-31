package service

import (
	"testing"

	"app/internal/model"
)

func TestCutiServiceValidation(t *testing.T) {
	svc, err := NewCutiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	// invalid date order
	err = svc.Simpan(&model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: "2026-10-10", TanggalSelesai: "2026-10-01", JumlahHari: 1, Status: "draft"})
	if err != ErrCutiInvalid {
		t.Fatalf("expected invalid, got %v", err)
	}

	// invalid jenis
	err = svc.Simpan(&model.Cuti{KaryawanId: 1, Jenis: "xxx", TanggalMulai: "2026-10-01", TanggalSelesai: "2026-10-02", JumlahHari: 2, Status: "draft"})
	if err != ErrCutiInvalid {
		t.Fatalf("expected invalid jenis, got %v", err)
	}

	// valid create
	c := &model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: "2026-10-01", TanggalSelesai: "2026-10-02", JumlahHari: 2, Status: "submitted"}
	if err := svc.Simpan(c); err != nil {
		t.Fatalf("simpan: %v", err)
	}

	// overlapping submitted
	err = svc.Simpan(&model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: "2026-10-02", TanggalSelesai: "2026-10-03", JumlahHari: 2, Status: "submitted"})
	if err != ErrCutiOverlap {
		t.Fatalf("expected overlap, got %v", err)
	}

	// non-overlapping allowed
	c2 := &model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: "2026-10-10", TanggalSelesai: "2026-10-11", JumlahHari: 2, Status: "submitted"}
	if err := svc.Simpan(c2); err != nil {
		t.Fatalf("simpan2: %v", err)
	}

	// karyawan lain boleh tanggal sama
	if err := svc.Simpan(&model.Cuti{KaryawanId: 2, Jenis: "sakit", TanggalMulai: "2026-10-10", TanggalSelesai: "2026-10-10", JumlahHari: 1, Status: "submitted"}); err != nil {
		t.Fatalf("simpan karyawan lain: %v", err)
	}
}

func TestCutiServiceDateValidation(t *testing.T) {
	svc, err := NewCutiService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	for _, tanggal := range []string{"01-10-2026", "2026-10-32", "2026-02-30", "besok"} {
		err := svc.Simpan(&model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: tanggal, TanggalSelesai: "2026-12-31", JumlahHari: 1, Status: "submitted"})
		if err != ErrCutiInvalid {
			t.Fatalf("tanggal mulai %q: expected cuti invalid, got %v", tanggal, err)
		}
		err = svc.Simpan(&model.Cuti{KaryawanId: 1, Jenis: "tahunan", TanggalMulai: "2026-01-01", TanggalSelesai: tanggal, JumlahHari: 1, Status: "submitted"})
		if err != ErrCutiInvalid {
			t.Fatalf("tanggal selesai %q: expected cuti invalid, got %v", tanggal, err)
		}
	}

	if _, err := svc.Update(999, 1, "tahunan", "01-10-2026", "2026-10-02", "", "submitted", 2); err != ErrCutiNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}