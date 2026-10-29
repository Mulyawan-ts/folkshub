package model

type Departemen struct {
	ID        int    `json:"id"`
	Kode      string `json:"kode"`
	Nama      string `json:"nama"`
	Kepala    string `json:"kepala"`
	Deskripsi string `json:"deskripsi"`
}

func (d *Departemen) TableName() string {
	return "departemen"
}
