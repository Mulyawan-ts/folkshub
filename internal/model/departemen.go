package model

import "time"

type Departemen struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Kode      string    `json:"kode" gorm:"not null;uniqueIndex"`
	Nama      string    `json:"nama" gorm:"not null"`
	Kepala    string    `json:"kepala"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (d *Departemen) TableName() string {
	return "departemen"
}
