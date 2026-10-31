package utils

import "time"

const IsoDate = "2006-01-02"
const IsoMonth = "2006-01"

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
