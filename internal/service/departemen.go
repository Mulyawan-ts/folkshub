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
	ErrDepartemenNotFound = errors.New("departemen tidak ditemukan")
)

type DepartemenService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Departemen
	nextID int
}

func NewDepartemenService(dbPath string) (*DepartemenService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &DepartemenService{
		path:   abs,
		daftar: make(map[int]*model.Departemen),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *DepartemenService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Departemen
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, d := range list {
		s.daftar[d.ID] = d
		if d.ID >= s.nextID {
			s.nextID = d.ID + 1
		}
	}

	return nil
}

func (s *DepartemenService) save() error {
	list := make([]*model.Departemen, 0, len(s.daftar))
	for _, d := range s.daftar {
		list = append(list, d)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *DepartemenService) Simpan(d *model.Departemen) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if d.ID == 0 {
		d.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[d.ID]; exists {
		return errors.New("departemen sudah ada")
	}

	s.daftar[d.ID] = d
	return s.save()
}

func (s *DepartemenService) Ambil(id int) (*model.Departemen, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	d, ok := s.daftar[id]
	if !ok {
		return nil, ErrDepartemenNotFound
	}

	return d, nil
}

func (s *DepartemenService) Semua() []*model.Departemen {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Departemen, 0, len(s.daftar))
	for _, d := range s.daftar {
		result = append(result, d)
	}

	return result
}

func (s *DepartemenService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrDepartemenNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *DepartemenService) Update(id int, kode, nama, kepala, deskripsi string) (*model.Departemen, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	d, ok := s.daftar[id]
	if !ok {
		return nil, ErrDepartemenNotFound
	}

	d.Kode = kode
	d.Nama = nama
	d.Kepala = kepala
	d.Deskripsi = deskripsi

	s.daftar[id] = d
	if err := s.save(); err != nil {
		return nil, err
	}

	return d, nil
}
