package main

import (
	"context"
	"fmt"

	"app/internal/model"
	"app/internal/service"
)

// App struct
type App struct {
	ctx          context.Context
	karyawanSvc  *service.KaryawanService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		karyawanSvc: service.NewKaryawanService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// SimpanKaryawan menyimpan data karyawan baru
func (a *App) SimpanKaryawan(nama, email, status string) (*model.Karyawan, error) {
	k := &model.Karyawan{
		Nama:   nama,
		Email:  email,
		Status: status,
	}

	if err := a.karyawanSvc.Simpan(k); err != nil {
		return nil, err
	}

	return k, nil
}

// AmbilKaryawan mengembalikan data karyawan berdasarkan ID
func (a *App) AmbilKaryawan(id int) (*model.Karyawan, error) {
	return a.karyawanSvc.Ambil(id)
}

// SemuaKaryawan mengembalikan semua data karyawan
func (a *App) SemuaKaryawan() []*model.Karyawan {
	return a.karyawanSvc.Semua()
}

// HapusKaryawan menghapus data karyawan berdasarkan ID
func (a *App) HapusKaryawan(id int) error {
	return a.karyawanSvc.Hapus(id)
}
