package model

import "time"

type Karyawan struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama      string    `json:"nama" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;uniqueIndex"`
	Status    string    `json:"status" gorm:"not null"` // "outsourcing", "tetap", "magang"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

const (
	StatusOutsourcing = "outsourcing"
	StatusTetap       = "tetap"
	StatusMagang      = "magang"
)

func (k *Karyawan) TableName() string {
	return "karyawan"
}
