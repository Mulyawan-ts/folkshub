package handler

import (
	"app/internal/model"
	"app/internal/service"
)

// Karyawan membungkus binding Karyawan untuk Wails. Method-nya dipromosikan ke
// App lewat embedding, sehingga nama binding tetap App.SimpanKaryawan dll.
type Karyawan struct {
	Svc *service.KaryawanService
}

func (h *Karyawan) SimpanKaryawan(nama, email, status string) (*model.Karyawan, error) {
	k := &model.Karyawan{
		Nama:   nama,
		Email:  email,
		Status: status,
	}

	if err := h.Svc.Simpan(k); err != nil {
		return nil, err
	}

	return k, nil
}

func (h *Karyawan) AmbilKaryawan(id int) (*model.Karyawan, error) {
	return h.Svc.Ambil(id)
}

func (h *Karyawan) SemuaKaryawan() ([]*model.Karyawan, error) {
	return h.Svc.Semua()
}

func (h *Karyawan) HapusKaryawan(id int) error {
	return h.Svc.Hapus(id)
}

func (h *Karyawan) UpdateKaryawan(id int, nama, email, status string) (*model.Karyawan, error) {
	return h.Svc.Update(id, nama, email, status)
}
