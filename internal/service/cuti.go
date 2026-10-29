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
	ErrCutiNotFound = errors.New("cuti tidak ditemukan")
)

type CutiService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Cuti
	nextID int
}

func NewCutiService(dbPath string) (*CutiService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &CutiService{
		path:   abs,
		daftar: make(map[int]*model.Cuti),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *CutiService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Cuti
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, c := range list {
		s.daftar[c.ID] = c
		if c.ID >= s.nextID {
			s.nextID = c.ID + 1
		}
	}

	return nil
}

func (s *CutiService) save() error {
	list := make([]*model.Cuti, 0, len(s.daftar))
	for _, c := range s.daftar {
		list = append(list, c)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *CutiService) Simpan(c *model.Cuti) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if c.ID == 0 {
		c.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[c.ID]; exists {
		return errors.New("cuti sudah ada")
	}

	s.daftar[c.ID] = c
	return s.save()
}

func (s *CutiService) Ambil(id int) (*model.Cuti, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, ok := s.daftar[id]
	if !ok {
		return nil, ErrCutiNotFound
	}

	return c, nil
}

func (s *CutiService) Semua() []*model.Cuti {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Cuti, 0, len(s.daftar))
	for _, c := range s.daftar {
		result = append(result, c)
	}

	return result
}

func (s *CutiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrCutiNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *CutiService) Update(id int, karyawanId int, jenis, tanggalMulai, tanggalSelesai, keterangan, status string, jumlahHari int) (*model.Cuti, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.daftar[id]
	if !ok {
		return nil, ErrCutiNotFound
	}

	c.KaryawanId = karyawanId
	c.Jenis = jenis
	c.TanggalMulai = tanggalMulai
	c.TanggalSelesai = tanggalSelesai
	c.JumlahHari = jumlahHari
	c.Keterangan = keterangan
	c.Status = status

	s.daftar[id] = c
	if err := s.save(); err != nil {
		return nil, err
	}

	return c, nil
}
