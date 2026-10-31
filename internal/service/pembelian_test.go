package service

import (
	"errors"
	"testing"

	"app/internal/model"
)

func newPembelianService(t *testing.T) *PembelianService {
	t.Helper()
	svc, err := NewPembelianService(":memory:")
	if err != nil {
		t.Fatalf("new service: %v", err)
	}
	t.Cleanup(func() { _ = svc.Close() })
	return svc
}

func pembelianValid() *model.Pembelian {
	return &model.Pembelian{
		Tanggal:     "2026-10-31",
		Barang:      "Kertas A4",
		Jenis:       model.JenisATK,
		Jumlah:      10,
		Satuan:      "rim",
		HargaSatuan: 55000,
		MetodeBayar: model.MetodeTunai,
		Vendor:      "Toko Sinar",
		Departemen:  "HRD",
		Status:      "diajukan",
	}
}

func TestPembelianSimpanMenghitungTotal(t *testing.T) {
	svc := newPembelianService(t)
	p := pembelianValid()

	if err := svc.Simpan(p); err != nil {
		t.Fatalf("simpan: %v", err)
	}
	if p.ID == 0 {
		t.Fatalf("expected id assigned")
	}
	if p.Total != 550000 {
		t.Fatalf("total = %d, want 550000 (10 x 55000)", p.Total)
	}

	got, err := svc.Ambil(p.ID)
	if err != nil || got.Total != 550000 || got.Barang != "Kertas A4" {
		t.Fatalf("ambil: %v / %+v", err, got)
	}
}

func TestPembelianMenolakInputTidakValid(t *testing.T) {
	svc := newPembelianService(t)

	kasus := []struct {
		nama string
		ubah func(*model.Pembelian)
	}{
		{"barang kosong", func(p *model.Pembelian) { p.Barang = "" }},
		{"tanggal kosong", func(p *model.Pembelian) { p.Tanggal = "" }},
		{"tanggal bukan kalender", func(p *model.Pembelian) { p.Tanggal = "2026-02-30" }},
		{"jenis asing", func(p *model.Pembelian) { p.Jenis = "entah" }},
		{"status asing", func(p *model.Pembelian) { p.Status = "entah" }},
		{"metode bayar kosong", func(p *model.Pembelian) { p.MetodeBayar = "" }},
		{"metode bayar asing", func(p *model.Pembelian) { p.MetodeBayar = "kripto" }},
		{"jumlah nol", func(p *model.Pembelian) { p.Jumlah = 0 }},
		{"harga negatif", func(p *model.Pembelian) { p.HargaSatuan = -1 }},
	}

	for _, k := range kasus {
		p := pembelianValid()
		k.ubah(p)
		if err := svc.Simpan(p); !errors.Is(err, ErrInvalidInput) {
			t.Fatalf("%s: err = %v, want ErrInvalidInput", k.nama, err)
		}
	}
}

func TestPembelianSimpanMenolakIDYangSudahAda(t *testing.T) {
	svc := newPembelianService(t)
	p := pembelianValid()
	if err := svc.Simpan(p); err != nil {
		t.Fatalf("simpan: %v", err)
	}

	lain := pembelianValid()
	lain.ID = p.ID
	if err := svc.Simpan(lain); !errors.Is(err, ErrPembelianExists) {
		t.Fatalf("err = %v, want ErrPembelianExists", err)
	}
}

func TestPembelianUpdateMenghitungUlangTotal(t *testing.T) {
	svc := newPembelianService(t)
	p := pembelianValid()
	if err := svc.Simpan(p); err != nil {
		t.Fatalf("simpan: %v", err)
	}

	ubah := pembelianValid()
	ubah.Barang = "Tinta Printer"
	ubah.Jenis = model.JenisPeralatan
	ubah.Jumlah = 3
	ubah.HargaSatuan = 120000
	ubah.Status = "dibeli"

	got, err := svc.Update(p.ID, ubah)
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if got.Total != 360000 {
		t.Fatalf("total = %d, want 360000 (3 x 120000)", got.Total)
	}
	if got.Status != "dibeli" || got.Barang != "Tinta Printer" {
		t.Fatalf("field tidak terupdate: %+v", got)
	}
}

func TestPembelianUpdateMenolakIDTidakDikenal(t *testing.T) {
	svc := newPembelianService(t)
	if _, err := svc.Update(999, pembelianValid()); !errors.Is(err, ErrPembelianNotFound) {
		t.Fatalf("err = %v, want ErrPembelianNotFound", err)
	}
}

func TestPembelianHapusDanSemua(t *testing.T) {
	svc := newPembelianService(t)

	a := pembelianValid()
	b := pembelianValid()
	b.Barang = "Pulpen"
	if err := svc.Simpan(a); err != nil {
		t.Fatalf("simpan a: %v", err)
	}
	if err := svc.Simpan(b); err != nil {
		t.Fatalf("simpan b: %v", err)
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 2 {
		t.Fatalf("semua: %v / %d", err, len(list))
	}

	if err := svc.Hapus(a.ID); err != nil {
		t.Fatalf("hapus: %v", err)
	}
	list, err = svc.Semua()
	if err != nil || len(list) != 1 || list[0].Barang != "Pulpen" {
		t.Fatalf("setelah hapus: %v / %+v", err, list)
	}

	if err := svc.Hapus(a.ID); !errors.Is(err, ErrPembelianNotFound) {
		t.Fatalf("hapus dua kali: err = %v, want ErrPembelianNotFound", err)
	}
}

func TestPembelianMetodeBayarIndependenDariJenis(t *testing.T) {
	svc := newPembelianService(t)

	for _, metode := range []string{model.MetodeTunai, model.MetodeNonTunai} {
		for _, jenis := range []string{model.JenisATK, model.JenisLangganan} {
			p := pembelianValid()
			p.MetodeBayar = metode
			p.Jenis = jenis
			if err := svc.Simpan(p); err != nil {
				t.Fatalf("metode=%s jenis=%s: %v", metode, jenis, err)
			}
		}
	}

	list, err := svc.Semua()
	if err != nil || len(list) != 4 {
		t.Fatalf("semua: %v / %d, want 4 kombinasi", err, len(list))
	}
}
