package model

type Jabatan struct {
	ID          int    `json:"id"`
	Kode        string `json:"kode"`
	Nama        string `json:"nama"`
	Departemen  string `json:"departemen"`
	GajiPokok   int    `json:"gaji_pokok"`
	Tunjangan   int    `json:"tunjangan"`
	Deskripsi   string `json:"deskripsi"`
}

func (j *Jabatan) TableName() string {
	return "jabatan"
}
