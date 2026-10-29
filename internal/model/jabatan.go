package model

import "time"

type Jabatan struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Kode       string    `json:"kode" gorm:"not null;uniqueIndex"`
	Nama       string    `json:"nama" gorm:"not null"`
	Departemen string    `json:"departemen" gorm:"index"`
	GajiPokok  int       `json:"gaji_pokok" gorm:"not null"`
	Tunjangan  int       `json:"tunjangan" gorm:"not null"`
	Deskripsi  string    `json:"deskripsi"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (j *Jabatan) TableName() string {
	return "jabatan"
}