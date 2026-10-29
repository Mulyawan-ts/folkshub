package service

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"app/internal/model"
)

var (
	ErrPelatihanNotFound = errors.New("pelatihan tidak ditemukan")
)

type PelatihanService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Pelatihan
	nextID int
}

func NewPelatihanService(dbPath string) (*PelatihanService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &PelatihanService{
		path:   abs,
		daftar: make(map[int]*model.Pelatihan),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *PelatihanService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Pelatihan
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, p := range list {
		s.daftar[p.ID] = p
		if p.ID >= s.nextID {
			s.nextID = p.ID + 1
		}
	}

	return nil
}

func (s *PelatihanService) save() error {
	list := make([]*model.Pelatihan, 0, len(s.daftar))
	for _, p := range s.daftar {
		list = append(list, p)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *PelatihanService) Simpan(p *model.Pelatihan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p.ID == 0 {
		p.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[p.ID]; exists {
		return errors.New("pelatihan sudah ada")
	}

	s.daftar[p.ID] = p
	return s.save()
}

func (s *PelatihanService) Ambil(id int) (*model.Pelatihan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.daftar[id]
	if !ok {
		return nil, ErrPelatihanNotFound
	}

	return p, nil
}

func (s *PelatihanService) Semua() []*model.Pelatihan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Pelatihan, 0, len(s.daftar))
	for _, p := range s.daftar {
		result = append(result, p)
	}

	return result
}

func (s *PelatihanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrPelatihanNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *PelatihanService) Update(id int, karyawanID int, nama, jenis, tanggalMulai, tanggalSelesai, lokasi, sertifikat, catatan, status string, biaya int) (*model.Pelatihan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.daftar[id]
	if !ok {
		return nil, ErrPelatihanNotFound
	}

	p.KaryawanID = karyawanID
	p.Nama = nama
	p.Jenis = jenis
	p.TanggalMulai = tanggalMulai
	p.TanggalSelesai = tanggalSelesai
	p.Lokasi = lokasi
	p.Biaya = biaya
	p.Sertifikat = sertifikat
	p.Catatan = catatan
	p.Status = status

	s.daftar[id] = p
	if err := s.save(); err != nil {
		return nil, err
	}

	return p, nil
}
