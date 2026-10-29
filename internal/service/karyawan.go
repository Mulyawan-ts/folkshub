package service

import (
	"errors"
	"sync"

	"app/internal/model"
)

var (
	ErrKaryawanNotFound = errors.New("karyawan tidak ditemukan")
	ErrKaryawanExists   = errors.New("karyawan sudah terdaftar")
)

type KaryawanService struct {
	mu     sync.RWMutex
	daftar map[int]*model.Karyawan
	nextID int
}

func NewKaryawanService() *KaryawanService {
	return &KaryawanService{
		daftar: make(map[int]*model.Karyawan),
		nextID: 1,
	}
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
	return nil
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
	return nil
}
