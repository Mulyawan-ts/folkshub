package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestErrorLoggerWritesToFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app.log")

	l, err := NewErrorLogger(path)
	if err != nil {
		t.Fatalf("new logger: %v", err)
	}
	defer l.Close()

	l.Error("gagal simpan", "entity", "karyawan", "id", 1)
	l.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !strings.Contains(string(data), "gagal simpan") {
		t.Fatalf("log content missing: %s", data)
	}
}

func TestErrorLoggerRecoversPanic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app.log")

	l, err := NewErrorLogger(path)
	if err != nil {
		t.Fatalf("new logger: %v", err)
	}
	defer l.Close()

	func() {
		defer l.Recover("worker")
		panic("boom")
	}()

	l.Close()
	data, _ := os.ReadFile(path)
	if !strings.Contains(string(data), "boom") {
		t.Fatalf("panic not logged: %s", data)
	}
}
