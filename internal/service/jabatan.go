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
	ErrJabatanNotFound = errors.New("jabatan tidak ditemukan")
)

type JabatanService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Jabatan
	nextID int
}

func NewJabatanService(dbPath string) (*JabatanService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &JabatanService{
		path:   abs,
		daftar: make(map[int]*model.Jabatan),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *JabatanService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Jabatan
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, j := range list {
		s.daftar[j.ID] = j
		if j.ID >= s.nextID {
			s.nextID = j.ID + 1
		}
	}

	return nil
}

func (s *JabatanService) save() error {
	list := make([]*model.Jabatan, 0, len(s.daftar))
	for _, j := range s.daftar {
		list = append(list, j)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *JabatanService) Simpan(j *model.Jabatan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if j.ID == 0 {
		j.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[j.ID]; exists {
		return errors.New("jabatan sudah ada")
	}

	s.daftar[j.ID] = j
	return s.save()
}

func (s *JabatanService) Ambil(id int) (*model.Jabatan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	j, ok := s.daftar[id]
	if !ok {
		return nil, ErrJabatanNotFound
	}

	return j, nil
}

func (s *JabatanService) Semua() []*model.Jabatan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Jabatan, 0, len(s.daftar))
	for _, j := range s.daftar {
		result = append(result, j)
	}

	return result
}

func (s *JabatanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrJabatanNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *JabatanService) Update(id int, kode, nama, departemen, deskripsi string, gajiPokok, tunjangan int) (*model.Jabatan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	j, ok := s.daftar[id]
	if !ok {
		return nil, ErrJabatanNotFound
	}

	j.Kode = kode
	j.Nama = nama
	j.Departemen = departemen
	j.GajiPokok = gajiPokok
	j.Tunjangan = tunjangan
	j.Deskripsi = deskripsi

	s.daftar[id] = j
	if err := s.save(); err != nil {
		return nil, err
	}

	return j, nil
}
