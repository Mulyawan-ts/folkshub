package core

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
)

var ErrInvalidLogPath = errors.New("path log tidak boleh kosong")

// ErrorLogger menulis log error/crash ke file lokal.
type ErrorLogger struct {
	file   *os.File
	logger *slog.Logger
}

// NewErrorLogger membuka (atau membuat) file log dan menyiapkan slog JSON handler.
func NewErrorLogger(path string) (*ErrorLogger, error) {
	if path == "" {
		return nil, ErrInvalidLogPath
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &ErrorLogger{
		file:   f,
		logger: slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelWarn})),
	}, nil
}

// Error mencatat pesan error dengan atribut tambahan.
func (l *ErrorLogger) Error(msg string, attrs ...any) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Error(msg, attrs...)
}

// Recover menangkap panic, mencatat stack trace, dan menelan panic agar aplikasi tidak crash.
// Panggil via defer: defer l.Recover("konteks").
func (l *ErrorLogger) Recover(context string) {
	if r := recover(); r != nil {
		if l != nil && l.logger != nil {
			l.logger.Error("panic recovered",
				"context", context,
				"panic", fmt.Sprintf("%v", r),
				"stack", string(debug.Stack()),
			)
		}
	}
}

// Close menutup file log.
func (l *ErrorLogger) Close() error {
	if l == nil || l.file == nil {
		return nil
	}
	return l.file.Close()
}
