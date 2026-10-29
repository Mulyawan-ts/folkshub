package model

type Karyawan struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Status string `json:"status"` // "outsourcing", "tetap", "magang"
}

const (
	StatusOutsourcing = "outsourcing"
	StatusTetap       = "tetap"
	StatusMagang      = "magang"
)

func (k *Karyawan) TableName() string {
	return "karyawan"
}
