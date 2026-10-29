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
	ErrReimburseNotFound = errors.New("reimburse tidak ditemukan")
)

type ReimburseService struct {
	mu     sync.RWMutex
	path   string
	daftar map[int]*model.Reimburse
	nextID int
}

func NewReimburseService(dbPath string) (*ReimburseService, error) {
	abs, err := filepath.Abs(dbPath)
	if err != nil {
		return nil, err
	}

	s := &ReimburseService{
		path:   abs,
		daftar: make(map[int]*model.Reimburse),
		nextID: 1,
	}

	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

func (s *ReimburseService) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	var list []*model.Reimburse
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, r := range list {
		s.daftar[r.ID] = r
		if r.ID >= s.nextID {
			s.nextID = r.ID + 1
		}
	}

	return nil
}

func (s *ReimburseService) save() error {
	list := make([]*model.Reimburse, 0, len(s.daftar))
	for _, r := range s.daftar {
		list = append(list, r)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *ReimburseService) Simpan(r *model.Reimburse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r.ID == 0 {
		r.ID = s.nextID
		s.nextID++
	}

	if _, exists := s.daftar[r.ID]; exists {
		return errors.New("reimburse sudah ada")
	}

	s.daftar[r.ID] = r
	return s.save()
}

func (s *ReimburseService) Ambil(id int) (*model.Reimburse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	r, ok := s.daftar[id]
	if !ok {
		return nil, ErrReimburseNotFound
	}

	return r, nil
}

func (s *ReimburseService) Semua() []*model.Reimburse {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.Reimburse, 0, len(s.daftar))
	for _, r := range s.daftar {
		result = append(result, r)
	}

	return result
}

func (s *ReimburseService) Hapus(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.daftar[id]; !ok {
		return ErrReimburseNotFound
	}

	delete(s.daftar, id)
	return s.save()
}

func (s *ReimburseService) Update(id int, karyawanId int, tanggal, jenis, keterangan, bukti, status string, jumlah int) (*model.Reimburse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	r, ok := s.daftar[id]
	if !ok {
		return nil, ErrReimburseNotFound
	}

	r.KaryawanId = karyawanId
	r.Tanggal = tanggal
	r.Jenis = jenis
	r.Jumlah = jumlah
	r.Keterangan = keterangan
	r.Bukti = bukti
	r.Status = status

	s.daftar[id] = r
	if err := s.save(); err != nil {
		return nil, err
	}

	return r, nil
}
