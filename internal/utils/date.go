package utils

import "time"

const IsoDate = "2006-01-02"
const IsoMonth = "2006-01"

var hariID = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu",
}

func ValidDate(s string) bool {
	_, err := time.Parse(IsoDate, s)
	return err == nil
}

func ValidOptionalDate(s string) bool {
	return s == "" || ValidDate(s)
}

func ValidMonth(s string) bool {
	_, err := time.Parse(IsoMonth, s)
	return err == nil
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse(IsoDate, s)
}

func FormatDate(t time.Time) string {
	return t.Format(IsoDate)
}

// FormatID mengembalikan "Hari, DD-MM-YYYY" (contoh: "Senin, 31-10-2026").
func FormatID(t time.Time) string {
	return hariID[t.Weekday()] + ", " + t.Format("02-01-2006")
}

