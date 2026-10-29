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
	ErrKaryawanNotFound = errors.New("karyawan tidak ditemukan")
	ErrKaryawanExists   = errors.New("karyawan sudah terdaftar")
)

type KaryawanService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Karyawan
	nextID int
}

func NewKaryawanService(dbPath string) (*KaryawanService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &KaryawanService{
		path:   abs,
		daftar: make(map[int]*model.Karyawan),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *KaryawanService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Karyawan
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, k := range list {
		s.daftar[k.ID] = k
		if k.ID >= s.nextID {
			s.nextID = k.ID + 1
		}
	}

	return nil
}

func (s *KaryawanService) save() error {
	list := make([]*model.Karyawan, 0, len(s.daftar))
	for _, k := range s.daftar {
		list = append(list, k)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *KaryawanService) Simpan(k *model.Karyawan) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if k.ID == 0 {
		k.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[k.ID]; exists {
		return ErrKaryawanExists
	}

	s.daftar[k.ID] = k
	return s.save()
}

func (s *KaryawanService) Ambil(id int) (*model.Karyawan, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	k, ok := s.daftar[id]
	if !ok {
		return nil, ErrKaryawanNotFound
	}

	return k, nil
}

func (s *KaryawanService) Semua() []*model.Karyawan {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Karyawan, 0, len(s.daftar))
	for _, k := range s.daftar {
		result = append(result, k)
	}

	return result
}

func (s *KaryawanService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrKaryawanNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *KaryawanService) Update(id int, nama, email, status string) (*model.Karyawan, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	k, ok := s.daftar[id]
	if !ok {
		return nil, ErrKaryawanNotFound
	}

	k.Nama = nama
	k.Email = email
	k.Status = status

	s.daftar[id] = k
	if err := s.save(); err != nil {
		return nil, err
	}

	return k, nil
}
