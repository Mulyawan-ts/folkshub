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
	ErrGajiNotFound = errors.New("gaji tidak ditemukan")
)

type GajiService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Gaji
	nextID int
}

func NewGajiService(dbPath string) (*GajiService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &GajiService{
		path:   abs,
		daftar: make(map[int]*model.Gaji),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *GajiService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Gaji
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, g := range list {
		s.daftar[g.ID] = g
		if g.ID >= s.nextID {
			s.nextID = g.ID + 1
		}
	}

	return nil
}

func (s *GajiService) save() error {
	list := make([]*model.Gaji, 0, len(s.daftar))
	for _, g := range s.daftar {
		list = append(list, g)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *GajiService) Simpan(g *model.Gaji) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if g.ID == 0 {
		g.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[g.ID]; exists {
		return errors.New("gaji sudah ada")
	}

	g.Total = g.GajiPokok + g.Tunjangan - g.Potongan
	s.daftar[g.ID] = g
	return s.save()
}

func (s *GajiService) Ambil(id int) (*model.Gaji, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	g, ok := s.daftar[id]
	if !ok {
		return nil, ErrGajiNotFound
	}

	return g, nil
}

func (s *GajiService) Semua() []*model.Gaji {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Gaji, 0, len(s.daftar))
	for _, g := range s.daftar {
		result = append(result, g)
	}

	return result
}

func (s *GajiService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrGajiNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *GajiService) Update(id int, karyawanId int, bulan string, gajiPokok, tunjangan, potongan int, status string) (*model.Gaji, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	g, ok := s.daftar[id]
	if !ok {
		return nil, ErrGajiNotFound
	}

	g.KaryawanId = karyawanId
	g.Bulan = bulan
	g.GajiPokok = gajiPokok
	g.Tunjangan = tunjangan
	g.Potongan = potongan
	g.Total = gajiPokok + tunjangan - potongan
	g.Status = status

	s.daftar[id] = g
	if err := s.save(); err != nil {
		return nil, err
	}

	return g, nil
}
