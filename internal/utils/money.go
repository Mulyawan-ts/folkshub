package utils

import (
	"fmt"
	"strings"
)

// FormatRupiah memformat angka menjadi "Rp 1.234.567".
func FormatRupiah(n int) string {
	neg := n < 0
	if neg {
		n = -n
	}
	s := fmt.Sprintf("%d", n)
	// sisipkan titik setiap 3 digit dari kanan
	var b strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte('.')
		}
		b.WriteRune(c)
	}
	out := "Rp " + b.String()
	if neg {
		out = "-" + out
	}
	return out
}
