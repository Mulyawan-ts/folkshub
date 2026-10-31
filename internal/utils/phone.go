package utils

import "strings"

// NormalizePhone menghapus spasi/dash dan mengubah awalan +62/62 menjadi 0.
func NormalizePhone(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	if strings.HasPrefix(s, "+62") {
		s = "0" + s[3:]
	} else if strings.HasPrefix(s, "62") && len(s) > 9 {
		s = "0" + s[2:]
	}
	return s
}

// InternationalPhone mengembalikan format +62xxxxxxxxxx.
func InternationalPhone(s string) string {
	s = NormalizePhone(s)
	if strings.HasPrefix(s, "0") {
		s = "+62" + s[1:]
	}
	return s
}

// ValidPhone memeriksa nomor HP Indonesia sederhana (10-13 digit, mulai 08).
func ValidPhone(s string) bool {
	s = NormalizePhone(s)
	if len(s) < 10 || len(s) > 13 {
		return false
	}
	if !strings.HasPrefix(s, "08") {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
