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
	ErrAbsensiNotFound = errors.New("absensi tidak ditemukan")
)

type AbsensiService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Absensi
	nextID int
}

func NewAbsensiService(dbPath string) (*AbsensiService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &AbsensiService{
		path:   abs,
		daftar: make(map[int]*model.Absensi),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *AbsensiService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Absensi
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, a := range list {
		s.daftar[a.ID] = a
		if a.ID >= s.nextID {
			s.nextID = a.ID + 1
		}
	}

	return nil
}

func (s *AbsensiService) save() error {
	list := make([]*model.Absensi, 0, len(s.daftar))
	for _, a := range s.daftar {
		list = append(list, a)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *AbsensiService) Simpan(a *model.Absensi) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if a.ID == 0 {
		a.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[a.ID]; exists {
		return errors.New("absensi sudah ada")
	}

	s.daftar[a.ID] = a
	return s.save()
}

func (s *AbsensiService) Ambil(id int) (*model.Absensi, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	a, ok := s.daftar[id]
	if !ok {
		return nil, ErrAbsensiNotFound
	}

	return a, nil
}

func (s *AbsensiService) Semua() []*model.Absensi {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Absensi, 0, len(s.daftar))
	for _, a := range s.daftar {
		result = append(result, a)
	}

	return result
}

func (s *AbsensiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrAbsensiNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *AbsensiService) Update(id int, karyawanID int, tanggal, checkIn, checkOut, status, keterangan string) (*model.Absensi, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	a, ok := s.daftar[id]
	if !ok {
		return nil, ErrAbsensiNotFound
	}

	a.KaryawanID = karyawanID
	a.Tanggal = tanggal
	a.CheckIn = checkIn
	a.CheckOut = checkOut
	a.Status = status
	a.Keterangan = keterangan

	s.daftar[id] = a
	if err := s.save(); err != nil {
		return nil, err
	}

	return a, nil
}
